package rdb_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/syncs/rdb"
)

func ExampleLocker() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.11:6379",
		Password: "",
		DB:       0,
	})
	locker := rdb.NewLocker(client)

	locker.Lock("test", func() error {
		// do something
		return nil
	})
}
