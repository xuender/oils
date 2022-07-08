package syncs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
)

func ExampleConsumer() {
	consumer := syncs.NewConsumer[int](10, 3, time.Millisecond)
	consumer.Consume(3, func(produceChan chan int) {
		for i := 0; i < 100; i++ {
			produceChan <- i
		}
	}, func(num int, elems []int) {
		time.Sleep(time.Microsecond)
		fmt.Println(num, elems)
	})
	// Output:
	// 1
}

func TestNewConsumer(t *testing.T) {
	t.Parallel()

	sum := 0
	consumer := syncs.NewConsumer[int](10, 3, time.Millisecond)
	consumer.Consume(3, func(produceChan chan int) {
		for i := 0; i < 100; i++ {
			produceChan <- i
		}
	}, func(num int, elems []int) {
		for _, elem := range elems {
			sum += elem
		}

		time.Sleep(time.Microsecond)
	})

	assert.Equal(t, 4950, sum)
}
