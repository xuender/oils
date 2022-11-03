package base_test

import (
	"fmt"
	"sort"

	"github.com/xuender/oils/base"
)

func ExampleSlice() {
	str := base.NewSlice("c", "a")

	sort.Sort(str)
	fmt.Println(str[0])
	fmt.Println(str[1])
	fmt.Println(str.All("a", "b"))
	fmt.Println(str.Any("a", "b"))
	fmt.Println(str.Has("b"))
	fmt.Println(str.Index("c"))
	fmt.Println(str.Index("b"))
	fmt.Println(str)

	// Output:
	// a
	// c
	// false
	// true
	// false
	// 1
	// -1
	// Slice[a c]
}
