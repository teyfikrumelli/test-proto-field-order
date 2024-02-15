package main

import (
	"fmt"
	"log"

	"github.com/teyfikrumelli/test-proto-field-order/out/messages/message_v1"
	"github.com/teyfikrumelli/test-proto-field-order/out/messages/message_v2"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a Version 1 message
	msgV1 := &message_v1.MyMessage{
		MyInt:    42,
		MyString: "Hello, world!",
	}

	// Serialize Version 1 message
	dataV1, err := proto.Marshal(msgV1)
	if err != nil {
		log.Fatalf("Marshaling error: %v", err)
	}

	// Simulate consuming Version 1 message with Version 2 schema
	var msgV2 message_v2.MyMessage
	if err := proto.Unmarshal(dataV1, &msgV2); err != nil {
		log.Fatalf("Unmarshaling error: %v", err)
	}

	fmt.Printf("V2 from V1 data: myString = %s, myInt = %d\n", msgV2.MyString, msgV2.MyInt)

	// Create a Version 2 message
	msgV2New := &message_v2.MyMessage{
		MyInt:    24,
		MyString: "Goodbye, world!",
	}

	// Serialize Version 2 message
	dataV2, err := proto.Marshal(msgV2New)
	if err != nil {
		log.Fatalf("Marshaling error: %v", err)
	}

	// Simulate consuming Version 2 message with Version 1 schema
	var msgV1FromV2 message_v1.MyMessage
	if err := proto.Unmarshal(dataV2, &msgV1FromV2); err != nil {
		log.Fatalf("Unmarshaling error: %v", err)
	}

	fmt.Printf("V1 from V2 data: myInt = %d, myString = %s\n", msgV1FromV2.MyInt, msgV1FromV2.MyString)
}
