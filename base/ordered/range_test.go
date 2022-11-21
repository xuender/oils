package ordered_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/ordered"
)

func TestRange(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []int{0, 1, 2, 3}, ordered.Range(4))
	assert.Equals(t, []int{1, 2, 3}, ordered.Range(1, 4))
	assert.Equals(t, []int{0, 2, 4}, ordered.Range(0, 6, 2))
	assert.Nil(t, ordered.Range(6, 3))
	assert.Nil(t, ordered.Range())
	assert.Equals(t, []int{6, 5, 4}, ordered.Range(6, 3, -1))
	assert.Equals(t, []int{0, -1, -2, -3}, ordered.Range(-4))
}
