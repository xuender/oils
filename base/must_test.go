package base_test

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestMust(t *testing.T) {
	t.Parallel()
	assert.Panics(t, func() {
		base.Must(io.ErrClosedPipe)
	})
}

func TestMust1(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Must1(func() (bool, error) {
			return false, os.ErrExist
		}())
	})
	assert.Panics(t, func() {
		base.Must1(strconv.ParseInt("a", 10, 64))
	})
}

func TestMust2(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Must2(func() (bool, int, error) {
			return false, 1, os.ErrExist
		}())
	})
}

func TestMust3(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		base.Must3(func() (bool, int, bool, error) {
			return false, 1, false, os.ErrExist
		}())
	})
}
