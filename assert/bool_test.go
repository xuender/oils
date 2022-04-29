package assert_test

import (
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestTrue(t *testing.T) {
	t.Parallel()

	True(t, True(&errorfer{}, true))
	False(t, True(&errorfer{}, false))
}

func TestFalse(t *testing.T) {
	t.Parallel()

	True(t, False(&errorfer{}, false))
	False(t, False(&errorfer{}, true))
}
