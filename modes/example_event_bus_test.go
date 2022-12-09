package modes_test

import (
	"fmt"
	"time"

	"github.com/xuender/oils/modes"
)

func ExampleEventBus() {
	bus := modes.NewEventBus[int, string]()
	ob1 := make(chan int)
	ob2 := make(chan int)

	go func() {
		for num := range ob1 {
			fmt.Println("ob1", num)
		}
	}()

	go func() {
		for num := range ob2 {
			fmt.Println("ob2", num)
		}
	}()

	bus.Register(ob1, []string{"click", "key"})
	bus.Register(ob2, []string{"touch", "click"})
	bus.Post(modes.NewEvent("startup", 0))
	bus.Post(modes.NewEvent("touch", 1))
	bus.Post(modes.NewEvent("click", 2))
	time.Sleep(time.Millisecond)
	bus.Post(modes.NewEvent("key", 3))

	close(ob1)
	time.Sleep(time.Millisecond)
	bus.Post(modes.NewEvent("click", 4))
	time.Sleep(time.Millisecond)

	// Output:
	// ob2 1
	// ob1 2
	// ob2 2
	// ob1 3
	// ob2 4
}
