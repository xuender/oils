package syncs_test

import (
	"context"
	"testing"
	"time"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs"
)

func TestBeakyBucket_Consume(t *testing.T) {
	t.Parallel()

	bucket := syncs.NewLeakyBucket(100, time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	count := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				bucket.Consume(1)
				count++
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	assert.GreaterOrEqual(t, count, 200)
	assert.LessOrEqual(t, count, 300)
}

func TestBeakyBucket_TryConsume(t *testing.T) {
	t.Parallel()

	bucket := syncs.NewLeakyBucket(100, time.Second)
	count := 0

	for {
		if bucket.TryConsume(1) {
			count++
		} else {
			break
		}
	}

	assert.Equal(t, 100, count)
}
