package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/oils/syncs/rdb"
)

type Sub struct {
	id int
}

func (p *Sub) Update(count uint64) {
	fmt.Printf("id: %d, count: %d\n", p.id, count)
}

func (p *Sub) ID() int {
	return p.id
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.11:6379",
		Password: "",
		DB:       0,
	})
	ins := rdb.NewInstanceNum(client)
	sub := ins.Subject()
	sub.Register(&Sub{id: 1})
	sub.Register(&Sub{id: 2})

	go ins.Run()

	time.Sleep(10 * time.Second)
}
