package cache_test

import (
	"fmt"
	"time"

	"github.com/xuender/oils/dbs/cache"
)

func ExampleCache() {
	data := cache.NewCache[string, string](3 * time.Second)

	_ = data.Set("key1", "value1")
	_ = data.SetByDuration("key2", "value2", time.Second)
	_ = data.Set("key3", "value3")

	fmt.Println(data.Get("key1"))
	fmt.Println("init size:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println(data.Get("key3")) // reset expire time.
	fmt.Println("2 second:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println("4 second:", data.Size())

	// Output:
	// value1 true
	// init size: 3
	// value3 true
	// 2 second: 2
	// 4 second: 1
}

func ExampleNewCallbackCache() {
	data := cache.NewCallbackCache(3*time.Second, func(key, value string) {
		fmt.Println("del:", key, value)
	})

	_ = data.Set("key1", "value1")
	_ = data.SetByDuration("key2", "value2", time.Second)
	_ = data.Set("key3", "value3")

	fmt.Println(data.Get("key1"))
	fmt.Println("init size:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println(data.Get("key3")) // reset expire time.
	fmt.Println("2 second:", data.Size())
	time.Sleep(2 * time.Second)
	fmt.Println("4 second:", data.Size())

	// Output:
	// value1 true
	// init size: 3
	// value3 true
	// del: key2 value2
	// 2 second: 2
	// del: key1 value1
	// 4 second: 1
}
