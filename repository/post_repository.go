package repository

import "github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"

type PostRepository interface {
	Save(post entity.Post) ([]entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

func (repo *repo) Save(post entity.Post) ([]entity.Post, error) {
	posts := []entity.Post{{Id: 1, Title: "Title 1", Text: "Text 1"}, {Id: 2, Title: "Title 2", Text: "Text 2"}}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	return posts, nil
}

func (repo *repo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	posts = []entity.Post{{Id: 1, Title: "Title 1", Text: "Text 1"}, {Id: 2, Title: "Title 2", Text: "Text 2"}}
	return posts, nil
}
