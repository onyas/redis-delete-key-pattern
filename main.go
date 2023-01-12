package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var c Config

func init() {
	c.getConfig()
}

func main() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{c.Redis.Host},
	})

	for _, searchPattern := range c.Redis.DeletedKeys {
		var count int = 0
		ctx := context.Background()
		iter := rdb.Scan(ctx, 0, searchPattern, 0).Iterator()
		fmt.Printf("YOUR SEARCH PATTERN= %s\n", searchPattern)
		for iter.Next(ctx) {
			key := iter.Val()
			go func(key string) {
				rdb.Del(ctx, key)
				fmt.Printf("Deleted= %s\n", key)
			}(key)
			count++
		}
		if err := iter.Err(); err != nil {
			panic(err)
		}

		fmt.Printf("Deleted Count %d\n", count)
	}

	rdb.Close()
}
