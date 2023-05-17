package repository

import "github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"

type PostRepository interface {
	Save(post entity.Post) ([]entity.Post, error)
	FindAll() ([]entity.Post, error)
}
