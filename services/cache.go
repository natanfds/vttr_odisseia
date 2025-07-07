package services

import (
	"time"

	"github.com/go-redis/redis"
)

type cacheService struct {
	rdb *redis.Client
}

func (c *cacheService) Start() error {
	c.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := c.rdb.Ping().Err()
	return err
}

func (c *cacheService) Get(key string) (string, error) {
	val, err := c.rdb.Get(key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (c *cacheService) Set(key string, value string, ttl time.Duration) error {
	return c.rdb.Set(key, value, ttl).Err()
}

func (c *cacheService) Delete(key string) error {
	return c.rdb.Del(key).Err()
}

var CacheService cacheService = cacheService{}
