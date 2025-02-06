package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type DB struct {
	client *redis.Client
}

func New(config Config) DB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	return DB{client: rdb}
}

func (d DB) Client() *redis.Client {
	return d.client
}

// SetToken stores the JWT token inside the redis cache store
func (d DB) SetToken(ctx context.Context, userID uint, token string, expTime time.Duration) error {
	_, err := d.client.Set(ctx, fmt.Sprintf("userID:%d", userID), token, expTime).Result()
	if err != nil {
		return fmt.Errorf("could not store the token in the cache store: %v", err)
	}

	return nil
}

// GetToken returns the JWT token stored in the redis and tells if the token exists or not
func (d DB) GetToken(ctx context.Context, userID uint) (string, bool) {
	token, err := d.client.Get(ctx, fmt.Sprintf("userID:%d", userID)).Result()
	if err != nil {
		return "", false
	}

	return token, true
}

// DeleteToken removes the user's token from the redis
func (d DB) DeleteToken(ctx context.Context, userID uint) error {
	_, err := d.client.Del(ctx, fmt.Sprintf("userID:%d", userID)).Result()
	if err != nil {
		return fmt.Errorf("could not delete the token from cache store: %v", err)
	}

	return nil
}
