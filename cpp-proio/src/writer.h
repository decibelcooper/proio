#ifndef PROIO_WRITER_H
#define PROIO_WRITER_H

#include <string>

#include <google/protobuf/io/zero_copy_stream_impl.h>

#include "event.h"
#include "proio.pb.h"

namespace proio {
typedef proto::BucketHeader_CompType Compression;
const Compression LZ4 = proto::BucketHeader::LZ4;
const Compression GZIP = proto::BucketHeader::GZIP;
const Compression UNCOMPRESSED = proto::BucketHeader::NONE;

const std::size_t minBucketWriteWindow = 0x100000;

class BucketOutputStream : public google::protobuf::io::ZeroCopyOutputStream {
   public:
    BucketOutputStream() { offset = 0; }
    virtual ~BucketOutputStream() { ; }

    bool Next(void **data, int *size) {
        if (bytes.size() - offset < minBucketWriteWindow) bytes.resize(offset + minBucketWriteWindow);
        *data = &bytes[offset];
        *size = bytes.size() - offset;
        offset = bytes.size();
        return true;
    }
    void BackUp(int count) {
        offset -= count;
        if (offset < 0) offset = 0;
    }
    google::protobuf::int64 ByteCount() const { return offset; }
    bool AllowsAliasing() { return false; }

    uint8_t *Bytes() { return &bytes[0]; }
    void Reset() { offset = 0; }

   private:
    std::vector<uint8_t> bytes;
    uint64_t offset;
};

class Writer {
   public:
    Writer(int fd);
    Writer(std::string filename);
    virtual ~Writer();

    void Flush();
    void Push(Event *event);
    void SetCompression(Compression comp);

   private:
    void initBucket();

    BucketOutputStream *bucket;
    google::protobuf::io::FileOutputStream *fileStream;

    uint64_t bucketEvents;
    Compression compression;
};

const uint64_t bucketDumpSize = 0x1000000;

const uint8_t magicBytes[] = {0xe1, 0xc1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                              0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00};

const class SerializationError : public std::exception {
    virtual const char *what() const throw() { return "Failed to serialize message"; }
} serializationError;

const class FileCreationError : public std::exception {
    virtual const char *what() const throw() { return "Failed to creating file for writing"; }
} fileCreationError;
}  // namespace proio

#endif  // PROIO_WRITER_H
