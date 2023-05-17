package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

var ctx = context.Background()

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Get(key string) (*entity.Post, error) {
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var post entity.Post
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post, nil
}

func (cache *redisCache) Set(key string, value *entity.Post) error {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(ctx, key, json, cache.expires*time.Second)

	return nil
}
