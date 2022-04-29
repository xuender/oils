package assert_test

import (
	"errors"
	"testing"

	"github.com/xuender/oils/assert"
)

func TestPanics(t *testing.T) {
	t.Parallel()

	assert.True(t, assert.Panics(&errorfer{}, func() {
		panic("error")
	}))
	assert.False(t, assert.Panics(&errorfer{}, func() {}))
}

var errTest = errors.New("error")

func TestPanicsWithError(t *testing.T) {
	t.Parallel()
	assert.True(t, assert.PanicsWithError(&errorfer{}, "error", func() {
		panic("error")
	}))
	assert.True(t, assert.PanicsWithError(&errorfer{}, "error", func() {
		panic(errTest)
	}))
	assert.True(t, assert.PanicsWithError(&errorfer{}, "1", func() {
		panic(1)
	}))
	assert.False(t, assert.PanicsWithError(&errorfer{}, "error", func() {}))
	assert.False(t, assert.PanicsWithError(&errorfer{}, "error", func() {
		panic("other")
	}))
}
