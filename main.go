package main

import (
	"context"
	"log"

	"github.com/joelrose/etcd-redis/etcd"
	"github.com/joelrose/etcd-redis/redis"
)

func main() {
	etcdEndpoint := "http://localhost:2379"
	redisAddress := "localhost:6379"

	cache1, err := etcd.New(etcdEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer cache1.Close()

	cache2 := redis.New(redisAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer cache2.Close()

	ctx, cancel := context.WithCancel(context.Background())

	err = cache1.Set(ctx, "test", "12345")
	if err != nil {
		log.Fatalf("cannot set value %v", err)
	}

	err = cache2.Set(ctx, "test", "12345")
	if err != nil {
		log.Fatalf("cannot set value cache2 %v", err)
	}

	cancel()
}
