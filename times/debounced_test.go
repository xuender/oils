package times_test

import (
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/times"
)

func TestDebounced(t *testing.T) {
	t.Parallel()

	count := 0
	debouncedFn := times.Debounced(func() {
		count++
	}, 10*time.Millisecond)

	for i := 0; i < 100; i++ {
		debouncedFn()
	}

	time.Sleep(20 * time.Millisecond)

	assert.Equal(t, 1, count)
}
