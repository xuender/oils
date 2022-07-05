package dbs_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/dbs"
)

func TestInt2Byte(t *testing.T) {
	t.Parallel()

	num, bytes := dbs.Int2Byte(0)

	assert.Equal(t, 0, num)
	assert.Equal(t, 1, bytes)

	num, bytes = dbs.Int2Byte(1)

	assert.Equal(t, 0, num)
	assert.Equal(t, 2, bytes)

	num, bytes = dbs.Int2Byte(10)

	assert.Equal(t, 1, num)
	assert.Equal(t, 4, bytes)
}

func TestByte2Int(t *testing.T) {
	t.Parallel()

	num1, bs1 := dbs.Int2Byte(1)

	assert.Equals(t, []int{1}, dbs.Byte2Ints(num1, bs1))

	num2, bs2 := dbs.Int2Byte(4)

	assert.Equals(t, []int{4}, dbs.Byte2Ints(num2, bs2))
	assert.Equals(t, []int{1, 4}, dbs.Byte2Ints(num2, bs1|bs2))
}
