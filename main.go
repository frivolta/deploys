package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	v := os.Getenv("ENV_VARIABLE_RANDOM")
	router.HandleFunc("/hello", helloWorld()).Methods("GET")
	fmt.Println("Running on port 8081 with start sh", v)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func helloWorld() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World with start sh"))
	}
}
