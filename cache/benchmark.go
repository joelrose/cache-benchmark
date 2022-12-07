package cache

import (
	"context"
	"log"
	"strconv"
	"time"
)

type Benchmark struct {
	ctx        context.Context
	cache      Cache
	cacheAlias string
}

const (
	operationCount = 10000
)

func NewBenchmark(ctx context.Context, cache Cache, cacheAlias string) *Benchmark {
	return &Benchmark{ctx, cache, cacheAlias}
}

func (b *Benchmark) Run() {
	b.cacheSet()
	b.cacheGet()
}

func (b *Benchmark) runWithTime(name string, fn func(context.Context, Cache)) {
	log.Printf("Starting %s.%s benchmark", b.cacheAlias, name)

	startTime := time.Now()

	fn(b.ctx, b.cache)

	elapsedTime := time.Since(startTime)
	log.Printf("Operation took: %s and averaged %v per request", elapsedTime, elapsedTime/operationCount)
}

func (b *Benchmark) cacheSet() {
	b.runWithTime("set", func(ctx context.Context, c Cache) {
		for i := 0; i < operationCount; i++ {
			err := b.cache.Set(b.ctx, strconv.Itoa(i), strconv.Itoa(i))
			if err != nil {
				log.Fatalf("One operation failed: %v", err)
			}
		}
	})
}

func (b *Benchmark) cacheGet() {
	b.runWithTime("get", func(ctx context.Context, c Cache) {
		for i := 0; i < operationCount; i++ {
			_, err := c.Get(ctx, strconv.Itoa(i))
			if err != nil {
				log.Fatalf("One operation failed: %v", err)
			}
		}
	})
}
