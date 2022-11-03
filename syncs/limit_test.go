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
	limit := syncs.NewLimit(1000)

	for f := 0; f < 4; f++ {
		go func() {
			for i := 0; i < 500; i++ {
				limit.Wait()
			}
		}()
	}

	for i := 0; i < 2000; i++ {
		limit.Wait()
	}

	assert.GreaterOrEqual(t, times.ClockStop(start), time.Second*4)
}
