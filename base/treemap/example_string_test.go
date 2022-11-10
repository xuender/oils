package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/base/treemap"
)

func ExampleString() {
	tmap := treemap.NewString(-1)

	for i := 100; i < 200; i++ {
		tmap.Set(fmt.Sprintf("a%d", i), i)
	}

	for i := 0; i < 10; i++ {
		tmap.Set(fmt.Sprintf("b%d", i), i)
	}

	list := []int{}

	tmap.Prefix(func(key string, value int) bool {
		list = append(list, value)

		return true
	}, "b")

	fmt.Println(list)

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
