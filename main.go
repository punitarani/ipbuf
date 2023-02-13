package main

import (
	"io"
	"log"
	"net/http"
	"os"

	ipbuf "ipbuf/proto"

	"google.golang.org/protobuf/proto"
)

// main is the entry point for the program
func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	// Get user input string from HTTP POST request
	http.HandleFunc("/api/httptriggeripbuf", handler)

	// Start the HTTP server
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
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

		// Set the response headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/protobuf")

		// Send the serialized message to the client
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "OPTIONS" {
		// Handle CORS preflight requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,POST")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Return 204 No Content
		w.WriteHeader(http.StatusNoContent)
		return
	} else {
		// Handle non-POST requests
		http.Error(w, "Only POST requests are supported.", http.StatusMethodNotAllowed)
		return
	}
}
