package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	var myaws AWS

	// Open AWS session
	err := myaws.Open("us-east-1")
	if err != nil {
		log.Fatalf("Failed to open AWS session: %v", err)
	}
	defer myaws.Close()

	// Initialize AWS resources
	err = myaws.init()
	if err != nil {
		log.Fatalf("Failed to initialize AWS resources: %v", err)
	}

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
