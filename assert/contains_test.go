package assert_test

import (
	"testing"

	"github.com/xuender/oils/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Contains(&errorfer{}, []int{1, 2}, 1))
	assert.False(t, assert.Contains(&errorfer{}, []int{1, 2}, 3))
}
