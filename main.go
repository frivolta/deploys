package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld()).Methods("GET")
	fmt.Println("Running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func helloWorld() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	}
}
