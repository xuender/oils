package syncs_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/syncs"
)

func TestLeakyBucket_Consume(t *testing.T) {
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

func TestLeakyBucket_TryConsume(t *testing.T) {
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

func TestLeakyBucket_Consume2(t *testing.T) {
	t.Parallel()

	bucket := syncs.NewLeakyBucket(3, time.Second)
	bucket.Consume(5)

	assert.True(t, bucket.TryConsume(1))
}

func TestNewLeakyBucket(t *testing.T) {
	t.Parallel()

	bucket := syncs.NewLeakyBucket(100, time.Second)

	assert.NotNil(t, bucket)
}
