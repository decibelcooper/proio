
package proio;

import java.io.*;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.google.protobuf.ByteString;
import com.google.protobuf.Descriptors;
import com.google.protobuf.Message;

public class Seek {
    public static void main(String[] args) {
        long index=Long.valueOf(args[1]);
        if (args.length != 1) { 
            System.out.println("Please provide one argument for the input file name");
            return;
          }
        try {
          event_seeker(args[0], index);
        } catch (IOException e) {
          e.printStackTrace();
        }
        return;
       
    }
  
    private static String event_seeker(String args, Long index) throws IOException {

         try{
            Reader reader = new Reader(args);
          reader.skip(index);
          Event event = reader.next();
          for (long entryID : event.getTaggedEntries("MCParameters")) {
            System.out.println(event.getEntry(entryID).toString());
          }
      } catch (Throwable e) {
        System.out.println(e);
      }
      return "";
    }
  }
  