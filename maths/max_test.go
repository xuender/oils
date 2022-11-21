package maths_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/maths"
)

func TestMax(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, maths.Max(1, 2, 3))
	assert.Equal(t, -3, maths.Max(-11, -20, -3))
}
