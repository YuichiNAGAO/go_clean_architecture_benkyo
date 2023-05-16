package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/controller"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/infrastructure"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/initializer"

	gorilla "github.com/gorilla/mux"
)

func init() {
	// .env fileをloadし、環境変数として扱えるようにする
	initializer.LoadEnvVariables()
	infrastructure.ConnectDB()
}

func main() {
	route := gorilla.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}).Methods("GET")
	route.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	route.HandleFunc("/posts", controller.AddPost).Methods("POST")
	// curl -v -X POST -H "Content-Type: application/json" -d '{"id":1, "title":"タイトル", "text":"本文"}'  localhost:8080/posts

	log.Println("Server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", route))
}
