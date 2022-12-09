package modes_test

import (
	"fmt"
	"time"

	"github.com/xuender/oils/modes"
)

func ExampleSubject() {
	sub := modes.NewSubject[string]()
	update := make(chan string)

	go func() {
		for str := range update {
			fmt.Println(str)
		}
	}()

	sub.Register(update)

	sub.Notify("test1")
	sub.Notify("test2")
	time.Sleep(time.Millisecond)
	close(update)
	sub.Notify("test3")

	// Output:
	// test1
	// test2
}
