package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"url-shortener/models"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache() *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",
		DB:       0,             
	})

	return &RedisCache{Client: rdb}
}

func (r *RedisCache) Get(key string) (*models.URLData, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var urlData models.URLData
	err = json.Unmarshal([]byte(val), &urlData)
	if err != nil {
		return nil, err
	}
	return &urlData, nil
}

func (r *RedisCache) Set(key string, value models.URLData, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisCache) CountClick(alias string, maxLimit int, ttl time.Duration) error {
	key := "click:" + alias

	cnt, err := r.Client.Get(ctx, key).Int()
	if err == redis.Nil {
		return r.Client.Set(ctx, key, 1, ttl).Err()
	} else if err != nil {
		return err
	}

	if cnt >= maxLimit {
		return fmt.Errorf("click Limit Reached. Try after 2 mins")
	}

	return r.Client.Incr(ctx, key).Err()
}
    