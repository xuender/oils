package syncs_test

import (
	"fmt"
	"time"

	"github.com/xuender/oils/syncs"
)

func ExampleConsumeFunc() {
	syncs.ConsumeFunc(10, 3, 3, time.Millisecond, func(produceChan chan int) {
		for i := 0; i < 100; i++ {
			produceChan <- i
		}
	}, func(num int, elems []int) {
		time.Sleep(time.Microsecond)
		fmt.Println(num, elems)
	})
}
