package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type cache struct {
	client *redis.Client
}

func (c *cache) Set(key string, value interface{}, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (c *cache) Get(key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}

func New() *cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &cache{client: client}
}
