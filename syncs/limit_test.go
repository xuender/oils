package syncs_test

import (
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
	"github.com/xuender/oils/times"
)

func TestNewLimit(t *testing.T) {
	t.Parallel()

	start := times.ClockStart()
	limit := syncs.NewLimit(100)

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 50; i++ {
				limit.Wait()
			}
		}()
	}

	for i := 0; i < 200; i++ {
		limit.Wait()
	}

	assert.GreaterOrEqual(t, times.ClockStop(start), time.Second*4)
}

func TestLimit_Try(t *testing.T) {
	t.Parallel()

	limit := syncs.NewLimit(1000)
	assert.NotNil(t, limit.Try())
	time.Sleep(time.Millisecond)
	assert.Nil(t, limit.Try())
}
