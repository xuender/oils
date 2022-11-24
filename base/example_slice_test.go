package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleSlice() {
	str := base.NewSlice("c", "a")

	fmt.Println(str[0])
	fmt.Println(str[1])
	fmt.Println(str)

	// Output:
	// c
	// a
	// Slice[c a]
}
