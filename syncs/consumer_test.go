package syncs_test

import (
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
)

func TestNewConsumerFunc(t *testing.T) {
	t.Parallel()

	sum := 0

	syncs.ConsumeFunc(10, 3, 3, time.Millisecond, func(produceChan chan int) {
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

	pac := &testPAC{sum: 0}
	count := syncs.Consume[int](10, 3, 3, time.Millisecond, pac)

	assert.Equal(t, 4950, pac.sum)
	assert.Equal(t, 100, count)
}
