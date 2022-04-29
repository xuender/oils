package assert_test

import (
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestNil(t *testing.T) {
	t.Parallel()

	True(t, Nil(&errorfer{}, nil))
	False(t, Nil(&errorfer{}, 1))
}

func TestNotNil(t *testing.T) {
	t.Parallel()

	True(t, NotNil(&errorfer{}, 1))
	False(t, NotNil(&errorfer{}, nil))
}
