package proio;

import java.io.*;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.google.protobuf.ByteString;
import com.google.protobuf.Descriptors;
import com.google.protobuf.Message;
import net.sourceforge.argparse4j.ArgumentParsers;
import net.sourceforge.argparse4j.inf.ArgumentParser;
import net.sourceforge.argparse4j.inf.Namespace;
import java.util.ArrayList;

public class Info {

    public static void main(String[] args) {

        try {
            ArgumentParser parser = ArgumentParsers.newFor("java proio.Info").build()
                    .description("Output metadata for proio stream.");
            parser.addArgument("files")
                    .type(String.class)
                    .nargs("+")
                    .help("input a file or files");
            Namespace res = parser.parseArgs(args);
            ArrayList<String> files = res.get("files");

            for (String file : files) {
                if (files.size() > 1) {
                    System.out.println("Filename: " + file);
                }
                Reader reader = new Reader(file);
                Event storeevent = null;
                Event tmpEvent;
                while (true) {
                    tmpEvent = null;
                    try {
                        tmpEvent = reader.next(true);
                    } catch (Throwable e) {
                        ;
                    }
                    if (tmpEvent == null) {
                        break;
                    }
                    storeevent = tmpEvent;
                }
                Map<String, ByteString> metadata = storeevent.getMetadata();
                for (Map.Entry<String, ByteString> entry : metadata.entrySet()) {
                    if (entry.getValue().size() < 1000) {
                        System.out.println("Metadata: " + entry.getKey() + ": " + entry.getValue().toStringUtf8());
                    }
                }
            }
        } catch (Throwable e) {
            System.out.println(e);
        }
    }
}
