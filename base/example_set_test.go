package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleSet() {
	set := base.NewSet(1, 2, 2)

	set.Add(3)

	fmt.Println(len(set))
	fmt.Println(set.Any(7, 2))
	fmt.Println(set.All(7, 2))

	set.Del(1)

	fmt.Println(len(set))

	// Output:
	// 3
	// true
	// false
	// 2
}
