package service

import (
	"errors"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/repository"
)

type PostService interface {
	Validate(post entity.Post) error
	Create(post entity.Post) ([]entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo = repository.NewPostRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post entity.Post) error {
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (s *service) Create(post entity.Post) ([]entity.Post, error) {
	return repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
