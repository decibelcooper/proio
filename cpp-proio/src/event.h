#ifndef PROIO_EVENT_H
#define PROIO_EVENT_H

#include <string>

#include "proio.pb.h"

namespace proio {
/** Class representing a single event
 */
class Event {
   public:
    Event(proto::Event *eventProto = NULL);
    virtual ~Event();

    /** AddEntry takes a tag and protobuf message entry and adds it to the
     * Event.  The return value is a uint64_t ID number used to reference the
     * added entry.  Once the entry is added, it is owned by the Event object.
     */
    uint64_t AddEntry(google::protobuf::Message *entry, std::string tag = "");
    /** GetEntry takes an entry ID and returns the corresponding entry.  The
     * returned entries are still owned by the Event object.
     */
    google::protobuf::Message *GetEntry(uint64_t id);
    /** TagEntry adds a tag to an entry that has already been added, identified
     * by its ID.
     */
    void TagEntry(uint64_t id, std::string tag);
    /** UntagEntry removes a tag from a specified entry.
     */
    void UntagEntry(uint64_t id, std::string tag);
    /** RemoveEntry removes an entry from the event.
     */
    void RemoveEntry(uint64_t id);
    /** Tags returns a list of tags that exist in the event.
     */
    std::vector<std::string> Tags();
    /** TaggedEntries tages a tag string and returns a list of entry IDs that
     * the tag references.
     */
    std::vector<uint64_t> TaggedEntries(std::string tag);
    /** AllEntries returns IDs for all entries in the event.
     */
    std::vector<uint64_t> AllEntries();
    /** EntryTags performs a reverse lookup of tags that point to an entry.
     * This is a relatively expensive search.
     */
    std::vector<std::string> EntryTags(uint64_t id);
    /** DeleteTag removes a tag from the Event.
     */
    void DeleteTag(std::string tag);
    /** Metadata returns a mapping from a string key to a pointer to a string
     * that contains the metadata, by reference.  These metadata are all the
     * entries received on the stream up to this Event.
     */
    const std::map<std::string, std::shared_ptr<std::string>> &Metadata() { return metadata; }

    /** String returns a human-readable string representing the event.
     */
    std::string String();

   private:
    uint64_t getTypeID(google::protobuf::Message *entry);
    const google::protobuf::Descriptor *getDescriptor(uint64_t typeID);
    void tagCleanup();

    friend class Writer;
    void flushCache();
    proto::Event *getProto();

    proto::Event *eventProto;
    std::map<std::string, uint64_t> revTypeLookup;
    std::map<uint64_t, google::protobuf::Message *> entryCache;
    std::map<uint64_t, const google::protobuf::Descriptor *> descriptorCache;
    bool dirtyTags;

    friend class Reader;
    std::map<std::string, std::shared_ptr<std::string>> metadata;
};

const class UnknownMessageTypeError : public std::exception {
    virtual const char *what() const throw() { return "Unknown message type"; }
} unknownMessageTypeError;
}  // namespace proio

#endif  // PROIO_EVENT_H
