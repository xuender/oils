package maths_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/maths"
)

func TestRand(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		assert.NotEqual(t, maths.Rand(), maths.Rand())
		assert.Greater(t, uint64(1<<54), maths.Rand())
	}
}

func TestRandInt(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		num := maths.RandInt(1, 100)

		assert.GreaterOrEqual(t, num, 1)
		assert.Less(t, num, 100)
	}
}

func TestRandString(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		assert.Equal(t, 10, len(maths.RandString(10)))
		assert.NotEqual(t, maths.RandString(10), maths.RandString(10))
	}
}

func TestRandBytes(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		assert.Equal(t, 10, len(maths.RandBytes(10)))
		assert.NotEquals(t, maths.RandBytes(10), maths.RandBytes(10))
	}
}
