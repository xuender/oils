package maths_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/maths"
)

func TestMin(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, maths.Min(1, 2, 3))
	assert.Equal(t, -20, maths.Min(-11, -20, -3))
}
