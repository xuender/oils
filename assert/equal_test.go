package assert_test

import (
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	True(t, Equal(&errorfer{}, 1, 1))
	False(t, Equal(&errorfer{}, 1, 2))
}

func TestEquals(t *testing.T) {
	t.Parallel()

	True(t, Equals(&errorfer{}, []int{1}, []int{1}))
	False(t, Equals(&errorfer{}, []int{1}, []int{1, 2}))
	False(t, Equals(&errorfer{}, []int{1}, []int{2}))
}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	True(t, NotEqual(&errorfer{}, 1, 2))
	False(t, NotEqual(&errorfer{}, 1, 1))
}
