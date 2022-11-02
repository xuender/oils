package rdb_test

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/syncs/rdb"
	"github.com/xuender/oils/times"
)

func ExampleNewLimit() {
	clean := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.11:6379",
		Password: "",
		DB:       0,
	})

	limit := rdb.NewLimit(clean, "limit", 10)
	clock := times.ClockStart()

	for i := 0; i < 30; i++ {
		limit.Wait()
	}

	limit.Try()

	fmt.Println(times.ClockStop(clock))
}
