package main

import (
	"encoding/json"
	"net/http"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
)

var (
	posts []entity.Post
)

func init() {
	posts = []entity.Post{{Id: 1, Title: "Title 1", Text: "Text 1"}, {Id: 2, Title: "Title 2", Text: "Text 2"}}
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

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Error unmarshalling the request"}`))
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal(posts)
	w.Write(result)
}
