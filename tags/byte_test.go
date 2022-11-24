package tags_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/tags"
)

func TestInt2Byte(t *testing.T) {
	t.Parallel()

	num, bytes := tags.Int2Byte(0)

	assert.Equal(t, 0, num)
	assert.Equal(t, byte(1), bytes)

	num, bytes = tags.Int2Byte(1)

	assert.Equal(t, 0, num)
	assert.Equal(t, byte(2), bytes)

	num, bytes = tags.Int2Byte(10)

	assert.Equal(t, 1, num)
	assert.Equal(t, byte(4), bytes)
}

func TestByte2Int(t *testing.T) {
	t.Parallel()

	num1, bs1 := tags.Int2Byte(1)

	assert.Equal(t, []int{1}, tags.Byte2Ints(num1, bs1))

	num2, bs2 := tags.Int2Byte(4)

	assert.Equal(t, []int{4}, tags.Byte2Ints(num2, bs2))
	assert.Equal(t, []int{1, 4}, tags.Byte2Ints(num2, bs1|bs2))
}
