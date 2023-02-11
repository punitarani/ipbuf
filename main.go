package main

import (
	"io"
	"net/http"

	ipbuf "ipbuf/proto"

	"google.golang.org/protobuf/proto"
)

// main is the entry point for the program
func main() {
	// Get user input string from HTTP POST request
	http.HandleFunc("/", handler)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// handler is the base HTTP handler for the program
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get the POST request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Limit the size of the POST request body to 1MB
		if len(body) > (1024 * 1024) {
			http.Error(w, "Request body too large.", http.StatusRequestEntityTooLarge)
			return
		}

		// Build the IPBuf protobuf struct
		message := ipbuf.IPBuf{Msg: string(body)}

		// Serialize the message to protobuf
		data, err := proto.Marshal(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send the serialized message to the client
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Handle non-POST requests
		http.Error(w, "Only POST requests are supported.", http.StatusMethodNotAllowed)
		return
	}
}
