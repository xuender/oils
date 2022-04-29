package assert_test

import (
	"errors"
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestPanics(t *testing.T) {
	t.Parallel()

	True(t, Panics(&errorfer{}, func() {
		panic("error")
	}))
	False(t, Panics(&errorfer{}, func() {}))
}

var errTest = errors.New("error")

func TestPanicsWithError(t *testing.T) {
	t.Parallel()
	True(t, PanicsWithError(&errorfer{}, "error", func() {
		panic("error")
	}))
	True(t, PanicsWithError(&errorfer{}, "error", func() {
		panic(errTest)
	}))
	True(t, PanicsWithError(&errorfer{}, "1", func() {
		panic(1)
	}))
	False(t, PanicsWithError(&errorfer{}, "error", func() {}))
	False(t, PanicsWithError(&errorfer{}, "error", func() {
		panic("other")
	}))
}
