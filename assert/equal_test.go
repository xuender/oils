package assert_test

import (
	"testing"

	"github.com/xuender/oils/assert"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Equal(&errorfer{}, 1, 1))
	assert.False(t, assert.Equal(&errorfer{}, 1, 2))
}

func TestEquals(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Equals(&errorfer{}, []int{1}, []int{1}))
	assert.False(t, assert.Equals(&errorfer{}, []int{1}, []int{1, 2}))
	assert.False(t, assert.Equals(&errorfer{}, []int{1}, []int{2}))
}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.NotEqual(&errorfer{}, 1, 2))
	assert.False(t, assert.NotEqual(&errorfer{}, 1, 1))
}

func TestNotEquals(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.NotEquals(&errorfer{}, []int{1}, []int{2}))
	assert.True(t, assert.NotEquals(&errorfer{}, []int{1}, []int{2, 3}))
	assert.False(t, assert.NotEquals(&errorfer{}, []int{1}, []int{1}))
}
