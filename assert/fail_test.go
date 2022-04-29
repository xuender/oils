package assert_test

import (
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestMessage(t *testing.T) {
	t.Parallel()

	Equal(t, "", Message())
	Equal(t, "1", Message(1))
	Equal(t, "1(2)", Message(1, 2))
	Equal(t, "1(2, 3)", Message(1, 2, 3))
	Equal(t, "1(2, 3, \"x\")", Message(1, 2, 3, "x"))
}

func TestFail(t *testing.T) {
	t.Parallel()

	False(t, Fail(&errorfer{}, "fail"))
	False(t, Fail(&errorfer{}, "fail", 1, 2, 3))
}
