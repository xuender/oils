package ordered_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/ordered"
)

func TestIndex(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, ordered.Index([]int{1, 2, 3}, 2))
	assert.Equal(t, 2, ordered.Index([]int{1, 2, 3}, 3))
	assert.Equal(t, 0, ordered.Index([]int{1, 2, 3}, 1))
	assert.Equal(t, -1, ordered.Index([]int{1, 2, 3}, 4))
	assert.Equal(t, 6, ordered.Index([]int{1, 2, 3, 4, 5, 6, 7, 8}, 7))
}
