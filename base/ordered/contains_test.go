package ordered_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/ordered"
)

func TestContains(t *testing.T) {
	t.Parallel()

	assert.True(t, ordered.Contains([]int{1, 2, 3}, 2))
	assert.False(t, ordered.Contains([]int{1, 2, 3}, 4))
}
