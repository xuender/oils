package rdb_test

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/syncs/rdb"
)

// nolint: testableexamples
func ExampleLimits() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.11:6379",
		Password: "",
		DB:       0,
	})

	ins := rdb.NewInstanceNum(client)
	limits := rdb.NewLimits()

	ins.Register(limits.Update())

	go ins.Run()

	start := time.Now()
	limit := limits.QPS("test", 1000)

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

	fmt.Println(time.Since(start))
}
