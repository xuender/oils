package assert_test

import (
	"testing"

	. "github.com/xuender/oils/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()

	True(t, Contains(&errorfer{}, []int{1, 2}, 1))
	False(t, Contains(&errorfer{}, []int{1, 2}, 3))
}
