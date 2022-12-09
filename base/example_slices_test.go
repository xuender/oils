package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleDel() {
	fmt.Println(base.Del([]int{1, 2, 3}))
	fmt.Println(base.Del([]int{}, 3))
	fmt.Println(base.Del([]int{1, 2, 3}, 3))
	fmt.Println(base.Del([]int{1, 2, 3, 4, 5}, 1, -2, 4))

	// Output:
	// [1 2 3]
	// []
	// [1 2]
	// [2 3 5]
}

func ExampleDelAll() {
	fmt.Println(base.DelAll([]int{1, 2, 3}))
	fmt.Println(base.DelAll([]int{}, 3))
	fmt.Println(base.DelAll([]int{1, 3, 2, 3}, 3))

	// Output:
	// [1 2 3]
	// []
	// [1 2]
}
