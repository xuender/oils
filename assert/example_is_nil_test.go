package assert_test

import (
	"fmt"

	"github.com/xuender/oils/assert"
)

func ExampleIsNil() {
	fmt.Println(assert.IsNil(1))
	fmt.Println(assert.IsNil(nil))

	// Output:
	// false
	// true
}
