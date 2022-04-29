package base_test

import (
	"os"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func test1(param int) (int, error) {
	return param, os.ErrExist
}

func test2(param1 int, param2 string) (int, string, error) {
	return param1, param2, os.ErrExist
}

func test3(param1 int, param2 string, param3 float32) (int, string, float32, error) {
	return param1, param2, param3, os.ErrExist
}

func TestPass1(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, base.Pass1(test1(3)))
}

func TestPass2(t *testing.T) {
	t.Parallel()

	ret1, ret2 := base.Pass2(test2(3, "3"))
	assert.Equal(t, 3, ret1)
	assert.Equal(t, "3", ret2)
}

func TestPass3(t *testing.T) {
	t.Parallel()

	ret1, ret2, ret3 := base.Pass3(test3(3, "3", 3.14))
	assert.Equal(t, 3, ret1)
	assert.Equal(t, "3", ret2)
	assert.Equal(t, 3.14, ret3)
}

func TestPass(t *testing.T) {
	t.Parallel()

	base.Pass(os.ErrClosed)
}
