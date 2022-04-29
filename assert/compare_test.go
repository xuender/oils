package assert_test

import (
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestGreater(t *testing.T) {
	t.Parallel()

	True(t, Greater(&errorfer{}, 2, 1))
	False(t, Greater(&errorfer{}, 1, 2))
	False(t, Greater(&errorfer{}, 1, 1))
}

func TestGreaterOrEqual(t *testing.T) {
	t.Parallel()

	True(t, GreaterOrEqual(&errorfer{}, 2, 1))
	False(t, GreaterOrEqual(&errorfer{}, 1, 2))
	True(t, GreaterOrEqual(&errorfer{}, 1, 1))
}

func TestLess(t *testing.T) {
	t.Parallel()

	True(t, Less(&errorfer{}, 1, 2))
	False(t, Less(&errorfer{}, 2, 1))
	False(t, Less(&errorfer{}, 1, 1))
}

func TestLessOrEqual(t *testing.T) {
	t.Parallel()

	True(t, LessOrEqual(&errorfer{}, 1, 2))
	False(t, LessOrEqual(&errorfer{}, 2, 1))
	True(t, LessOrEqual(&errorfer{}, 1, 1))
}
