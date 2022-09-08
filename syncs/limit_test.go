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
	cha := syncs.NewLimit(3)

	for i := 0; i < 5; i++ {
		cha.Try()
	}

	assert.GreaterOrEqual(t, times.ClockStop(start), time.Second)
}
