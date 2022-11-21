package syncs_test

import (
	"sync"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
)

func TestNewLimitQueue(t *testing.T) {
	t.Parallel()

	count := 0
	limit := syncs.NewLimitQueue(100, time.Second*3, func(i int) { count++ })

	limit.Update(100)
	limit.ID()

	group := sync.WaitGroup{}
	group.Add(4)

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 50; i++ {
				_ = limit.Add(i)
			}
			group.Done()
		}()
	}

	group.Wait()
	time.Sleep(time.Second * 3)

	assert.Equal(t, 200, count)

	for i := 0; i < 200; i++ {
		_ = limit.Add(i)
	}

	time.Sleep(time.Second * 3)
	assert.Equal(t, 400, count)
}

func TestNewLimitQueueAdd(t *testing.T) {
	t.Parallel()

	count := 0
	limit := syncs.NewLimitQueue(10, time.Second*3, func(i int) {
		count++
	})

	for i := 0; i < 30; i++ {
		_ = limit.Add(i)
	}

	limit.Update(100)

	for i := 0; i < 300; i++ {
		_ = limit.Add(i)
	}

	time.Sleep(time.Second * 4)
	assert.Equal(t, count, 330)
}
