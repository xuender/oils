package tags_test

import (
	"fmt"

	"github.com/xuender/oils/tags"
)

func ExampleBitMap() {
	tag := tags.NewBitMap(1, 2, 33)

	fmt.Println(tag.Has(1))
	fmt.Println(tag.Has(3))
	fmt.Println(tag.Has(33))
	fmt.Println(tag)

	// Output:
	// true
	// false
	// true
	// 1, 2, 33
}
