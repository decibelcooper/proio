#include <fcntl.h>
#include <sys/types.h>

#include <google/protobuf/io/gzip_stream.h>
#include <lz4.h>

#include "reader.h"
#include "writer.h"

using namespace proio;
using namespace google::protobuf;

Reader::Reader(int fd) {
    this->fd = fd;
    fileStream = new io::FileInputStream(fd);
    closeFDOnDelete = false;

    initBucket();
}

Reader::Reader(std::string filename) {
    fd = open(filename.c_str(), O_RDONLY);
    if (fd == -1) throw fileOpenError;
    fileStream = new io::FileInputStream(fd);
    closeFDOnDelete = true;

    initBucket();
}

Reader::~Reader() {
    pthread_mutex_lock(&mutex);
    pthread_mutex_unlock(&mutex);
    pthread_mutex_destroy(&mutex);

    if (bucketHeader) delete bucketHeader;
    delete compBucket;
    LZ4F_freeDecompressionContext(dctxPtr);
    delete bucket;
    delete fileStream;
    if (closeFDOnDelete) close(fd);
}

Event *Reader::Next() {
    pthread_mutex_lock(&mutex);
    Event *event = readFromBucket();
    pthread_mutex_unlock(&mutex);
    return event;
}

proto::BucketHeader *Reader::NextHeader() {
    pthread_mutex_lock(&mutex);
    readBucket(std::numeric_limits<uint64_t>::max());
    pthread_mutex_unlock(&mutex);
    return bucketHeader;
}

uint64_t Reader::Skip(uint64_t nEvents) {
    pthread_mutex_lock(&mutex);

    uint64_t bucketEventsLeft = 0;
    if (bucketHeader) {
        bucketEventsLeft = bucketHeader->nevents() - bucketEventsRead;
    }
    uint64_t nSkipped = 0;
    if (nEvents > bucketEventsLeft) {
        nSkipped += bucketEventsLeft;
        uint64_t n;
        while ((n = readBucket(nEvents - nSkipped)) > 0) nSkipped += n;
    }

    uint64_t nRead0 = bucketEventsRead;
    while (nSkipped < nEvents) {
        readFromBucket(false);
        if (bucketEventsRead == nRead0) break;
        nRead0 = bucketEventsRead;
        nSkipped++;
    }

    pthread_mutex_unlock(&mutex);
    return nSkipped;
}

void Reader::SeekToStart() {
    pthread_mutex_lock(&mutex);

    delete fileStream;
    if (lseek(fd, 0, SEEK_SET) == -1) throw seekError;
    fileStream = new io::FileInputStream(fd);

    readBucket();

    pthread_mutex_unlock(&mutex);
}

void Reader::initBucket() {
    compBucket = new BucketInputStream(0);
    bucketEventsRead = 0;
    bucketHeader = NULL;
    LZ4F_createDecompressionContext(&dctxPtr, LZ4F_VERSION);
    bucket = new BucketInputStream(0);

    pthread_mutex_init(&mutex, NULL);
}

Event *Reader::readFromBucket(bool doMerge) {
    if (bucket->BytesRemaining() == 0) readBucket();
    auto stream = new io::CodedInputStream(bucket);

    uint32_t protoSize;
    if (!stream->ReadLittleEndian32(&protoSize)) {
        delete stream;
        return NULL;
    }

    bucketEventsRead++;
    if (doMerge) {
        auto eventLimit = stream->PushLimit(protoSize);
        auto eventProto = new proto::Event;
        if (!eventProto->MergeFromCodedStream(stream) || !stream->ConsumedEntireMessage())
            throw deserializationError;
        stream->PopLimit(eventLimit);
        delete stream;
        auto event = new Event(eventProto);
        event->metadata = metadata;
        return event;
    } else {
        if (!stream->Skip(protoSize)) throw corruptBucketError;
        delete stream;
        return NULL;
    }
}

uint64_t Reader::readBucket(uint64_t maxSkipEvents) {
    auto stream = new io::CodedInputStream(fileStream);
    syncToMagic(stream);

    bucketEventsRead = 0;
    compBucket->Reset(0);
    bucket->Reset(0);

    uint32_t headerSize;
    if (!stream->ReadLittleEndian32(&headerSize)) {
        delete stream;
        return 0;
    }

    auto headerLimit = stream->PushLimit(headerSize);
    if (bucketHeader) delete bucketHeader;
    bucketHeader = new proto::BucketHeader;
    if (!bucketHeader->MergeFromCodedStream(stream) || !stream->ConsumedEntireMessage())
        throw deserializationError;
    stream->PopLimit(headerLimit);

    // Set metadata for future events
    for (auto keyValuePair : bucketHeader->metadata())
        metadata[keyValuePair.first] = std::make_shared<std::string>(keyValuePair.second);

    uint64_t bucketSize = bucketHeader->bucketsize();
    if (bucketHeader->nevents() > maxSkipEvents) {
        compBucket->Reset(bucketSize);
        if (!stream->ReadRaw(compBucket->Bytes(), bucketSize)) throw corruptBucketError;
        delete stream;
    } else {
        if (!stream->Skip(bucketSize)) throw corruptBucketError;
        delete stream;
        return bucketHeader->nevents();
    }

    switch (bucketHeader->compression()) {
        case LZ4: {
            bucket->Reset(dctxPtr, compBucket);
            break;
        }
        case GZIP: {
            io::GzipInputStream *gzipStream = new io::GzipInputStream(compBucket);
            bucket->Reset(*gzipStream);
            delete gzipStream;
            break;
        }
        default:
            BucketInputStream *tmpBucket = bucket;
            bucket = compBucket;
            compBucket = tmpBucket;
    }

    return 0;
}

uint64_t Reader::syncToMagic(io::CodedInputStream *stream) {
    uint8_t num;
    uint64_t nRead = 0;

    while (stream->ReadRaw(&num, 1)) {
        nRead++;

        if (num == magicBytes[0]) {
            bool goodSeq = true;

            for (int i = 1; i < 16; i++) {
                if (!stream->ReadRaw(&num, 1)) break;
                nRead++;

                if (num != magicBytes[i]) {
                    goodSeq = false;
                    break;
                }
            }
            if (goodSeq) break;
        }
    }
    return nRead;
}

BucketInputStream::BucketInputStream(uint64_t size) {
    offset = 0;
    bytes.resize(size);
    this->size = size;
}

BucketInputStream::~BucketInputStream() { ; }

inline bool BucketInputStream::Next(const void **data, int *size) {
    *data = &bytes[offset];
    *size = this->size - offset;
    offset = this->size;
    if (*size == 0) return false;
    return true;
}

inline void BucketInputStream::BackUp(int count) { offset -= count; }

inline bool BucketInputStream::Skip(int count) {
    offset += count;
    if (offset > size) {
        offset = size;
        return false;
    }
    return true;
}

inline int64 BucketInputStream::ByteCount() const { return offset; }

uint8_t *BucketInputStream::Bytes() { return &bytes[0]; }

uint64_t BucketInputStream::BytesRemaining() { return size - offset; }

void BucketInputStream::Reset(uint64_t size) {
    offset = 0;
    if (bytes.size() < size) bytes.resize(size);
    this->size = size;
}

uint64_t BucketInputStream::Reset(io::ZeroCopyInputStream &stream) {
    Reset(0);
    uint8_t *data;
    int size;
    while (stream.Next((const void **)&data, &size)) {
        offset = this->size;
        this->size += size;
        if (this->size > bytes.size()) bytes.resize(this->size);
        std::memcpy(&bytes[offset], data, size);
    }
    offset = 0;
    return this->size;
}

uint64_t BucketInputStream::Reset(LZ4F_dctx *dctxPtr, BucketInputStream *compBucket) {
    offset = 0;
    size = bytes.size();
    if (size == 0) Reset(minBucketWriteWindow);
    int srcSize;
    uint8_t *srcBuffer;
    compBucket->Next((const void **)&srcBuffer, &srcSize);
    int srcBytesRemaining = srcSize;
    int dstSize;
    uint8_t *dstBuffer;
    size_t hint;
    while (srcBytesRemaining > 0) {
        Next((const void **)&dstBuffer, &dstSize);
        size_t srcSizeTmp = srcSize;
        size_t dstSizeTmp = dstSize;
        hint = LZ4F_decompress(dctxPtr, dstBuffer, &dstSizeTmp, srcBuffer, &srcSizeTmp, NULL);
        if (LZ4F_isError(hint)) throw badLZ4FrameError;
        srcBytesRemaining -= srcSizeTmp;
        BackUp(dstSize - dstSizeTmp);
        if (offset == size) {
            size += minBucketWriteWindow;
            bytes.resize(size);
        }
        compBucket->BackUp(srcBytesRemaining);
        compBucket->Next((const void **)&srcBuffer, &srcSize);
    }
    size = offset;
    offset = 0;
    if (hint != 0) {
        LZ4F_resetDecompressionContext(dctxPtr);
        throw badLZ4FrameError;
    }
    return size;
}
