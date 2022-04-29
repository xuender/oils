package base_test

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestPanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Panic(io.ErrClosedPipe)
	})
}

func TestPanic1(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Panic1(func() (bool, error) {
			return false, os.ErrExist
		}())
	})
	assert.Panics(t, func() {
		base.Panic1(strconv.ParseInt("a", 10, 64))
	})
}

func TestPanic2(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Panic2(func() (bool, int, error) {
			return false, 1, os.ErrExist
		}())
	})
}

func TestPanic3(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Panic3(func() (bool, int, bool, error) {
			return false, 1, false, os.ErrExist
		}())
	})
}
