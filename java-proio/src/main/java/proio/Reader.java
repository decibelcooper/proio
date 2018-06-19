package proio;

import java.io.ByteArrayInputStream;
import java.io.FileInputStream;
import java.io.InputStream;
import java.io.IOException;
import java.lang.Iterable;
import java.util.Iterator;
import java.util.HashMap;
import java.util.Map;
import java.util.zip.GZIPInputStream;

import com.google.protobuf.ByteString;
import com.google.protobuf.CodedInputStream;
import com.google.protobuf.InvalidProtocolBufferException;
import net.jpountz.lz4.LZ4FrameInputStream;

public class Reader implements Iterable<Event>, Iterator<Event>
{
	private InputStream fileStream = null;
	private CodedInputStream stream = null;
    private CodedInputStream bucket = null;
    private Proto.BucketHeader bucketHeader = null;
    private int bucketEventsRead = 0;
    private int bucketIndex = 0;
    private Map<String, ByteString> metadata = new HashMap<String, ByteString>();

    private Event queuedEvent = null;

	public Reader(String filename)
			throws IOException {
		fileStream = new FileInputStream(filename);
		stream = CodedInputStream.newInstance(fileStream);
	}

	public Reader(InputStream fileStream)
			throws IOException {
		stream = CodedInputStream.newInstance(fileStream);
	}

    public Event next(boolean metaOnly)
            throws IOException {
        while (bucketHeader == null || bucketIndex >= bucketHeader.getNEvents()) {
            if (bucketHeader != null) {
                bucketIndex -= bucketHeader.getNEvents();
            }
            readHeader();
            if (bucketHeader == null) {
                return null;
            }
        }

        Event event = new Event();
        event.metadata = metadata;
        if (!metaOnly) {
            if (bucket == null) {
                readBucket();
            }
            try {
                readFromBucket(event);
            } catch (InvalidProtocolBufferException e) {
                readBucket();
                readFromBucket(event);
            }
        }

        return event;
    }

	public void close()
			throws IOException {
		if (fileStream != null) {
            fileStream.close();
        }
	}

	public Iterator<Event> iterator() {
		return this;
	}

	public boolean hasNext() {
        if (queuedEvent != null) {
            return true;
        }
        queuedEvent = next();
        if (queuedEvent != null) {
            return true;
        }
		return false;
	}

	public Event next() {
        if (queuedEvent != null) {
            Event returnEvent = queuedEvent;
            queuedEvent = null;
            return returnEvent;
        }
        try {
            return next(false);
        } catch (IOException e) {
            return null;
        }
	}

	public void remove() {}

    private void readHeader()
            throws IOException {
        bucket = null;
        bucketHeader = null;
        bucketEventsRead = 0;

        syncToMagic();

        int headerSize = stream.readRawLittleEndian32();

        int headerLimit = stream.pushLimit(headerSize);
        bucketHeader = Proto.BucketHeader.parseFrom(stream);
        stream.popLimit(headerLimit);

        for (Map.Entry<String, ByteString> entry : bucketHeader.getMetadataMap().entrySet()) {
            metadata.put(entry.getKey(), entry.getValue());
        }
    }

    private void readBucket()
            throws IOException {
        int bucketSize = (int) bucketHeader.getBucketSize();
        ByteArrayInputStream compBucket = new ByteArrayInputStream(stream.readRawBytes((int) bucketSize));

        switch (bucketHeader.getCompression()) {
            case GZIP:
                bucket = CodedInputStream.newInstance(new GZIPInputStream(compBucket));
                break;
            case LZ4:
                bucket = CodedInputStream.newInstance(new LZ4FrameInputStream(compBucket));
                break;
            default:
                bucket = CodedInputStream.newInstance(compBucket);
                break;
        }
    }

    private void readFromBucket(Event event)
            throws InvalidProtocolBufferException, IOException {
        while (bucketEventsRead <= bucketIndex) {
            int protoSize = bucket.readRawLittleEndian32();

            if (event != null && bucketEventsRead == bucketIndex) {
                int eventLimit = bucket.pushLimit(protoSize);
                event.eventProto = Proto.Event.parseFrom(bucket);
                bucket.popLimit(eventLimit);
            } else {
                bucket.skipRawBytes(protoSize);
            }

            bucketEventsRead++;
        }
        bucketIndex++;
    }

	private void syncToMagic()
			throws IOException {
		int nRead = 0;
        while (true) {
            byte thisByte = stream.readRawByte();

            if (thisByte == Writer.magicBytes[0]) {
                boolean goodSeq = true;

                for (int i = 1; i < 16; i++) {
                    thisByte = stream.readRawByte();

                    if (thisByte != Writer.magicBytes[i]) {
                        goodSeq = false;
                        break;
                    }
                }

                if (goodSeq) break;
            }
        }

		return;
	}
}
