package times

import (
	"sync"
	"time"
)

// nolint
var clocks = sync.Map{}

// ClockStart 计时器开始.
//
// Deprecated: Use time.Since instead.
func ClockStart() int64 {
	now := time.Now()
	clockID := now.UnixNano()

	for _, has := clocks.LoadOrStore(clockID, now); has; _, has = clocks.LoadOrStore(clockID, now) {
		clockID++
	}

	return clockID
}

// ClockCount 计次.
//
// Deprecated: Use time.Since instead.
func ClockCount(clock int64) time.Duration {
	now := time.Now()
	if old, has := clocks.LoadOrStore(clock, now); has {
		if t, ok := old.(time.Time); ok {
			return now.Sub(t)
		}
	}

	return time.Duration(0)
}

// ClockStop 计时器截止.
//
// Deprecated: Use time.Since instead.
func ClockStop(clock int64) time.Duration {
	now := time.Now()

	if old, has := clocks.LoadAndDelete(clock); has {
		if t, ok := old.(time.Time); ok {
			return now.Sub(t)
		}
	}

	return time.Duration(0)
}
