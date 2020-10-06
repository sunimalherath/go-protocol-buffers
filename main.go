package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/sunimalherath/protocol_buffer/protoc-01/src/simple/simplepb"
)

func main() {
	message := doSimple()

	readAndWriteToFile(message)
	readAndWriteToJSON(message)
}

func readAndWriteToFile(message proto.Message) {
	// 1. Write protocol buffer message to file
	if err := writeToFile("simple.bin", message); err != nil {
		fmt.Println("An error occurred while writig file", err)
	}

	// 2. Read file back to protocol buffer
	messageToRead := &simplepb.SimpleMessage{}
	if err := readFromFile("simple.bin", messageToRead); err != nil {
		fmt.Println("An error occurred while reading file", err)
	}
	// check if the message unmarshalled properly
	fmt.Println(messageToRead)

}

func readAndWriteToJSON(message proto.Message) {
	// Protocol Buffer Message -> JSON
	smAsString := toJSON(message)
	fmt.Println(smAsString)

	// JSON -> Protocol Buffer Message
	msgToReadFromJSON := &simplepb.SimpleMessage{}
	fromJSON(smAsString, msgToReadFromJSON)
	fmt.Println(msgToReadFromJSON)
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "Simple Message",
		SampleList: []int32{1, 2, 4, 6},
	}

	// fmt.Println(sm)

	// sm.Name = "Renamed Message"

	// fmt.Println(sm)
	// fmt.Println(sm.GetName())

	return &sm
}

// ------------ START - Read and Write protocol buffer message to file --------------------------

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cannot serialised to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fName, out, 0644); err != nil {
		log.Fatalln("Cannot write to file", err)
		return err
	}

	fmt.Println("Data written to file")
	return nil
}

func readFromFile(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("An error occurred while reading file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Could not put the bytes into the protocol buffers struct", err)
		return err
	}
	fmt.Println("Successfully read the file")
	return nil
}

// ------------ END - Read and Write protocol buffer message to file --------------------------

// ------------ START - Read and Write protocol buffer message to JSON ------------------------

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cannot convert to JSON", err)
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln("Cannot convert JSON to SimpleMessage struct", err)
	}

}

// ------------ END - Read and Write protocol buffer message to JSON ------------------------
