package cache

import "github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"

type PostCache interface {
	Get(key string) (*entity.Post, error)
	Set(key string, value *entity.Post) error
}
