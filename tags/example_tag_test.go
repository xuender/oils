package tags_test

import (
	"fmt"

	"github.com/xuender/oils/tags"
)

func ExampleTag() {
	var tag tags.Tag

	tag.Add(1, 2)
	fmt.Println(tag.Has(1))
	fmt.Println(tag.Has(3))

	tag.Add(3)
	fmt.Println(tag.Len())
	fmt.Println(tag.Slice())

	tag.Del(2)
	fmt.Println(tag.Len())

	fmt.Println(tag.All(3, 4))
	fmt.Println(tag.All(1, 3))
	fmt.Println(tag.Any(3, 4))
	fmt.Println(tag.Any(4, 5))

	// Output:
	// true
	// false
	// 3
	// [1 2 3]
	// 2
	// false
	// true
	// true
	// false
}
