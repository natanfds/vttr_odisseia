package services

import (
	"time"

	"github.com/go-redis/redis"

	"github.com/natanfds/vtt_odisseia/configs"
)

type cacheService struct {
	rdb *redis.Client
}

func (c *cacheService) Start() error {
	c.rdb = redis.NewClient(&redis.Options{
		Addr:     configs.ENV.RedisAddr,
		Password: configs.ENV.RedisPass,
		DB:       configs.ENV.RedisDb,
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
