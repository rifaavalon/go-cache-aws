package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	var myaws AWS

	myaws.Open("us-east-1")
	myaws.init()

	log.Fatal(http.ListenAndServe(":8080", router))
}

//HandleError for errors
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
