package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/joelrose/etcd-redis/cache"
	etcd "go.etcd.io/etcd/client/v3"
)

type Etcd struct {
	client *etcd.Client
}

func New(url string) (cache.Cache, error) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{url},
		DialTimeout: 1 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to etcd: %v", err)
	}

	return Etcd{client}, nil
}

func (e Etcd) Get(ctx context.Context, key string) (string, error) {
	resp, err := e.client.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return string(resp.Kvs[0].Value), nil
}

func (e Etcd) Set(ctx context.Context, key string, value string) error {
	_, err := e.client.Put(ctx, key, string(value))
	return err
}

func (e Etcd) Delete(ctx context.Context, key string) error {
	_, err := e.client.Delete(ctx, key)
	return err
}

func (e Etcd) Close() error {
	return e.client.Close()
}
