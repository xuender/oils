package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleSub() {
	fmt.Println(base.Sub([]int{1, 2, 3}, -1))
	fmt.Println(base.Sub([]int{1, 2, 3, 4, 5}, 1, -2))

	// Output:
	// [3]
	// [2 3]
}
