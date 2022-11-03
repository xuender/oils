package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestParseInteger(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 10, base.Must1(base.ParseInteger[int]("10")))
}

func TestParseInteger_Float(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 10, base.Must1(base.ParseInteger[int]("10.3")))
	assert.Equal(t, 11, base.Must1(base.ParseInteger[int]("10.5")))
}

func TestParseInteger_Error(t *testing.T) {
	t.Parallel()

	_, err := base.ParseInteger[int]("xxfef.3r3r")

	assert.NotNil(t, err)
}

func TestParseFloat(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3.14, base.Must1(base.ParseFloat[float64]("3.14")))
}

func TestItoa(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "3", base.Itoa(3))
	assert.Equal(t, "3", base.Itoa(3.14))
}

func TestFormatFloat(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "3", base.FormatFloat(3, 3))
	assert.Equal(t, "3.14", base.FormatFloat(3.14, 3))
}

func TestRound(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, base.Round[int](3.14))
	assert.Equal(t, 3, base.Round[int](2.74))
}

func TestParseIntegerAny(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int]("3")))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int]("3.0")))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](3)))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](3.0)))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](float32(3.0))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](uint(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](int8(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](uint8(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](int16(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](uint16(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](int32(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](uint32(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](int64(3))))
	assert.Equal(t, 3, base.Must1(base.ParseIntegerAny[int](uint64(3))))
	assert.Equal(t, 102, base.Must1(base.ParseIntegerAny[int]('f')))
	assert.Equal(t, 1000000, base.Must1(base.ParseIntegerAny[int]([]byte{0x40, 0x42, 0xf})))

	_, err := base.ParseIntegerAny[int]([]int{})
	assert.NotNil(t, err)
}
