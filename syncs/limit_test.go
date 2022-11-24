package syncs_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs"
)

func TestNewLimit(t *testing.T) {
	t.Parallel()

	start := time.Now()
	limit := syncs.NewLimit(1)

	limit.QPS(100)

	group := sync.WaitGroup{}
	group.Add(4)

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 50; i++ {
				limit.Wait()
			}
			group.Done()
		}()
	}

	group.Wait()

	assert.GreaterOrEqual(t, time.Since(start), time.Second*2)

	for i := 0; i < 200; i++ {
		limit.Wait()
	}

	assert.GreaterOrEqual(t, time.Since(start), time.Second*4)
}

func TestLimit_Try(t *testing.T) {
	t.Parallel()

	limit := syncs.NewLimit(1000)
	assert.NotNil(t, limit.Try())
	time.Sleep(time.Millisecond)
	assert.Nil(t, limit.Try())
}
