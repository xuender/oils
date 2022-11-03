package syncs_test

import (
	"fmt"
	"time"

	"github.com/xuender/oils/syncs"
)

func ExampleCounter() {
	counter := syncs.NewCounter[string]()

	counter.Inc("test1")
	counter.Inc("test2")
	counter.Inc("test1")
	counter.Add("test3", 5)
	counter.Dec("test3")

	fmt.Println(counter.Get("test1"))
	fmt.Println(counter.Get("test2"))
	fmt.Println(counter.Get("test3"))
	counter.Clean()
	fmt.Println(counter.Get("test1"))
	fmt.Println(counter.Size())

	// Output:
	// 2
	// 1
	// 4
	// 0
	// 1
}

func ExampleCounter_concurrent() {
	counter := syncs.NewCounter[int]()
	work := func() {
		for i := 0; i < 1000; i++ {
			counter.Inc(i)
		}
	}

	for i := 0; i < 1000; i++ {
		go work()
	}

	time.Sleep(time.Second)

	fmt.Println(counter.Sum())

	// Output:
	// 1000000
}
