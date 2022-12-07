package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/joelrose/etcd-redis/cache"
	"github.com/joelrose/etcd-redis/etcd"
	"github.com/joelrose/etcd-redis/redis"
)

func main() {
	etcdEndpoint := "http://localhost:2379"
	redisAddress := "localhost:6379"

	etcdClient, err := etcd.New(etcdEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer etcdClient.Close()

	redisClient := redis.New(redisAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	type test struct {
		name  string
		cache cache.Cache
	}

	tests := []test{
		{
			name:  "Redis",
			cache: redisClient,
		},
		{
			name:  "etcd",
			cache: etcdClient,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	for _, test := range tests {
		log.Printf("Starting Set test with %v", test.name)

		startTime := time.Now()

		for i := 0; i < 10000; i++ {
			err = test.cache.Set(ctx, strconv.Itoa(i), strconv.Itoa(i))
			if err != nil {
				log.Fatalf("One operation failed: %v", err)
			}
		}

		elapsedTime := time.Since(startTime)
		log.Printf("Operation took: %s", elapsedTime)
	}

	for _, test := range tests {
		log.Printf("Starting Get test with %v", test.name)

		startTime := time.Now()

		for i := 0; i < 10000; i++ {
			_, err := test.cache.Get(ctx, strconv.Itoa(i))
			if err != nil {
				log.Fatalf("One operation failed: %v", err)
			}
		}

		elapsedTime := time.Since(startTime)
		log.Printf("Operation took: %s", elapsedTime)
	}

	cancel()
}
