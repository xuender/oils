package ordered_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/ordered"
)

func ExampleTimes() {
	fmt.Println(ordered.Times(3, strconv.Itoa))

	// Output:
	// [0 1 2]
}

func TestTimes(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"0", "1", "2"}, ordered.Times(3, strconv.Itoa))
}
