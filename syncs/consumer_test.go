package syncs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/syncs"
)

func ExampleConsumer() {
	consumer := syncs.NewConsumer[int](10, 3, time.Millisecond)
	consumer.Consume(3, func(produceChan chan int) {
		for i := 0; i < 100; i++ {
			produceChan <- i
		}
	}, func(elems []int) {
		fmt.Println(elems)
	})
}

func TestNewConsumer(t *testing.T) {
	t.Parallel()

	set := base.NewSet[int]()
	consumer := syncs.NewConsumer[int](10, 3, time.Millisecond)
	consumer.Consume(3, func(produceChan chan int) {
		for i := 0; i < 100; i++ {
			produceChan <- i
		}
	}, func(elems []int) {
		set.Add(elems...)
	})

	assert.Equal(t, 100, len(set))
}
