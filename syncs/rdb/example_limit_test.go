package rdb_test

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/syncs/rdb"
	"github.com/xuender/oils/times"
)

func ExampleLimit() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.11:6379",
		Password: "",
		DB:       0,
	})

	limit := rdb.NewLimiter(client, "limit", 1000)
	clock := times.ClockStart()

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 500; i++ {
				limit.Wait()
			}
		}()
	}

	for i := 0; i < 2000; i++ {
		limit.Wait()
	}

	_ = limit.Try()

	fmt.Println(times.ClockStop(clock))
	// 1Output:
	// 4.188548639s
}
