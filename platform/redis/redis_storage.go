package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func (r *RedisStorage) Get(key string) ([]byte, error) {
	return r.client.Get(context.Background(), key).Bytes()
}

func (r *RedisStorage) Set(key string, value []byte, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisStorage) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisStorage) Reset() error {
	return r.client.FlushDB(context.Background()).Err()
}

func (r *RedisStorage) Close() error {
	return r.client.Close()
}

func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{client: client}
}
