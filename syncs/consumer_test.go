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
	consumer.ConsumeFunc(3, func(produceChan chan int) {
		for i := 0; i < 100; i++ {
			produceChan <- i
		}
	}, func(num int, elems []int) {
		time.Sleep(time.Microsecond)
		fmt.Println(num, elems)
	})
}

func TestNewConsumerFunc(t *testing.T) {
	t.Parallel()

	sum := 0
	consumer := syncs.NewConsumer[int](10, 3, time.Millisecond)
	consumer.ConsumeFunc(3, func(produceChan chan int) {
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

type testPAC struct {
	sum int
}

func (p *testPAC) Produc(produceChan chan int) {
	for i := 0; i < 100; i++ {
		produceChan <- i
	}
}

func (p *testPAC) Consum(num int, elems []int) {
	for _, elem := range elems {
		p.sum += elem
	}

	time.Sleep(time.Microsecond)
}

func TestNewConsumer(t *testing.T) {
	t.Parallel()

	pac := &testPAC{}
	chanConsumer := syncs.NewConsumer[int](10, 3, time.Millisecond)
	chanConsumer.Consume(3, pac)

	assert.Equal(t, 4950, pac.sum)
}
