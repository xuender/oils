package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestValues(t *testing.T) {
	t.Parallel()

	testMap := map[int]int{3: 1, 4: 6}
	assert.Equals(t, []int{6, 1}, base.Values(testMap, []int{4, 3}))
	assert.Equals(t, []int{6, 0}, base.Values(testMap, []int{4, 2}))
}
