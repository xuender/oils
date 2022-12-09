package modes_test

import (
	"fmt"
	"sort"
	"time"

	"github.com/xuender/oils/modes"
)

func ExampleEventBus() {
	bus := modes.NewEventBus[int, string]()
	ob1 := make(chan int)
	ob2 := make(chan int)
	out := []string{}

	go func() {
		for num := range ob1 {
			out = append(out, fmt.Sprintf("%d ob1", num))
		}
	}()

	go func() {
		for num := range ob2 {
			out = append(out, fmt.Sprintf("%d ob2", num))
		}
	}()

	fmt.Println(bus.Has("key"))
	bus.Register(ob1, []string{"click", "key"})
	fmt.Println(bus.Has("key"))
	bus.Register(ob2, []string{"touch", "click"})
	bus.Post(modes.NewEvent("startup", 0))
	bus.Post(modes.NewEvent("touch", 1))
	bus.Post(modes.NewEvent("click", 2))
	time.Sleep(time.Millisecond)
	bus.Post(modes.NewEvent("key", 3))

	close(ob1)
	time.Sleep(time.Millisecond)
	bus.Post(modes.NewEvent("key", 4))
	bus.Post(modes.NewEvent("click", 4))
	time.Sleep(time.Millisecond)
	fmt.Println(bus.Has("key"))

	sort.Strings(out)

	for _, str := range out {
		fmt.Println(str)
	}

	// Output:
	// false
	// true
	// false
	// 1 ob2
	// 2 ob1
	// 2 ob2
	// 3 ob1
	// 4 ob2
}
