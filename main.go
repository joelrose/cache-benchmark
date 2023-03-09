package main

import (
	"context"
	"log"

	"github.com/joelrose/etcd-redis/cache"
	"github.com/joelrose/etcd-redis/etcd"
	"github.com/joelrose/etcd-redis/redis"
)

const (
	OperationCount = 10000
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisBenchmark := cache.NewBenchmark(ctx, redisClient, "Redis")
	redisBenchmark.Run()

	etcdBenchmark := cache.NewBenchmark(ctx, etcdClient, "etcd")
	etcdBenchmark.Run()
}
