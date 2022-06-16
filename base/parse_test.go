package base_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func ExampleParseInteger() {
	// 字符串转换成整数
	fmt.Println(base.Must1(base.ParseInteger[int]("3")))

	// Output:
	// 3
}

func ExampleParseFloat() {
	fmt.Println(base.Must1(base.ParseFloat[float32]("3.14")))

	// Output:
	// 3.14
}

func ExampleItoa() {
	fmt.Println(base.Itoa(3))
	fmt.Println(base.Itoa(3.14))

	// Output:
	// 3
	// 3
}

func ExampleFormatFloat() {
	fmt.Println(base.FormatFloat(3, 3))
	fmt.Println(base.FormatFloat(3.14, 3))

	// Output:
	// 3
	// 3.14
}

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
