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
	route.HandleFunc("/posts", addPost).Methods("POST")
	// curl -v -X POST -H "Content-Type: application/json" -d '{"id":1, "title":"タイトル", "text":"本文"}'  localhost:8080/posts

	log.Println("Server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", route))
}
