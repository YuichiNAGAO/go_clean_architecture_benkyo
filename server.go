package main

import (
	"fmt"
	"log"
	"net/http"

	gorilla "github.com/gorilla/mux"
)

func main() {
	route := gorilla.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}).Methods("GET")
	route.HandleFunc("/posts", getPosts).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", route))
}
