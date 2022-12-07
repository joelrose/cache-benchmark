package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/joelrose/etcd-redis/cache"
)

type Redis struct {
	client *redis.Client
}

func New(address string) cache.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return Redis{client}
}

func (r Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	return val, err
}

func (r Redis) Set(ctx context.Context, key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r Redis) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r Redis) Close() error {
	return r.client.Close()
}
