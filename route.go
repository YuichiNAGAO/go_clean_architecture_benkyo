package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Text 1"}, {Id: 2, Title: "Title 2", Text: "Text 2"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Error marshalling the posts array"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
