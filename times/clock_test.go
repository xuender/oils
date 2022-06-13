package times_test

import (
	"strings"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/times"
)

func TestClockStop(t *testing.T) {
	t.Parallel()

	cid := times.ClockStart()

	time.Sleep(time.Second)

	t1 := times.ClockCount(cid)

	time.Sleep(time.Second)

	assert.True(t, strings.HasPrefix(times.Natural(t1), "1秒钟"))
	assert.True(t, strings.HasPrefix(times.Natural(times.ClockStop(cid)), "2秒钟"))
}

func TestClockStart(t *testing.T) {
	t.Parallel()

	cid1 := times.ClockStart()
	cid2 := times.ClockStart()

	times.ClockStop(cid1)
	times.ClockStop(cid2)

	assert.True(t, cid1 != cid2)
}

func TestClockCount(t *testing.T) {
	t.Parallel()

	assert.Equal(t, time.Duration(0), times.ClockStop(1))
	assert.Equal(t, time.Duration(0), times.ClockCount(1))
}
