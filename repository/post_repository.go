package repository

import (
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/infrastructure"
)

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

func (repo *repo) Save(post entity.Post) ([]entity.Post, error) {
	var posts []entity.Post
	infrastructure.DB.Find(&posts)
	infrastructure.DB.Create(&post)
	posts = append(posts, post)
	return posts, nil
}

func (repo *repo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	infrastructure.DB.Find(&posts)
	return posts, nil
}
