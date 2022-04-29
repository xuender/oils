package assert_test

import (
	"testing"

	"github.com/xuender/oils/assert"
)

func TestGreater(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Greater(&errorfer{}, 2, 1))
	assert.False(t, assert.Greater(&errorfer{}, 1, 2))
	assert.False(t, assert.Greater(&errorfer{}, 1, 1))
}

func TestGreaterOrEqual(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.GreaterOrEqual(&errorfer{}, 2, 1))
	assert.False(t, assert.GreaterOrEqual(&errorfer{}, 1, 2))
	assert.True(t, assert.GreaterOrEqual(&errorfer{}, 1, 1))
}

func TestLess(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Less(&errorfer{}, 1, 2))
	assert.False(t, assert.Less(&errorfer{}, 2, 1))
	assert.False(t, assert.Less(&errorfer{}, 1, 1))
}

func TestLessOrEqual(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.LessOrEqual(&errorfer{}, 1, 2))
	assert.False(t, assert.LessOrEqual(&errorfer{}, 2, 1))
	assert.True(t, assert.LessOrEqual(&errorfer{}, 1, 1))
}
