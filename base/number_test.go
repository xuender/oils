package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestJSNumber(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0x155dfb1875c8c7, base.JSNumber(0x5c955dfb1875c8c7))
	assert.Equal(t, 0x155dfb1875c8c7, base.JSNumber(0x5c955dfb1875c8c7))
}

// nolint
func FuzzJSNumber(f *testing.F) {
	f.Add(uint64(5))

	f.Fuzz(func(t *testing.T, number uint64) {
		t.Helper()

		assert.LessOrEqual(t, base.JSNumber(number), uint64(1<<53))
	})
}

func TestMax(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, base.Max(1, 2, 3))
	assert.Equal(t, -3, base.Max(-11, -20, -3))
}

func TestMin(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, base.Min(1, 2, 3))
	assert.Equal(t, -20, base.Min(-11, -20, -3))
}
