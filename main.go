package main

import (
	"fmt"

	ipbuf "ipbuf/proto"

	"google.golang.org/protobuf/proto"
)

// main is the entry point for the program
func main() {
	var input string

	// Get user input string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input: ", err)
	}

	// Build the IPBuf protobuf struct
	message := ipbuf.IPBuf{Msg: input}

	// Serialize the message to protobuf
	data, err := proto.Marshal(&message)
	fmt.Printf("Serialized data: %v. Size: %d bytes.\n", data, len(data))

	// Deserialize the message from protobuf
	var newMessage ipbuf.IPBuf
	err = proto.Unmarshal(data, &newMessage)
	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
	}
	fmt.Printf("Deserialized message: %v", newMessage.GetMsg())
}
