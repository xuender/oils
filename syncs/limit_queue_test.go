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

	limit.SetQPS(100)
	limit.SetTimeOut(time.Second * 3)
	assert.Equal(t, 100, limit.QPS())
	assert.Equal(t, time.Second*3, limit.TimeOut())

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
	assert.Equal(t, 0, limit.Len())
}

func TestNewLimitQueueAdd(t *testing.T) {
	t.Parallel()

	count := 0
	limit := syncs.NewLimitQueue(10, time.Second*2, func(i int) {
		count += i
	})

	for i := 0; i < 20; i++ {
		assert.Nil(t, limit.Add(1000))
	}

	time.Sleep(time.Millisecond)

	limit.SetQPS(100)

	for i := 0; i < 200; i++ {
		assert.Nil(t, limit.Add(1))
	}

	time.Sleep(time.Second * 3)
	assert.Equal(t, count, 20200)
}
