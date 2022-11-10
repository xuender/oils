package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/base/treemap"
)

func Example() {
	tmap := treemap.New(-1, -1)

	for i := 0; i < 10; i++ {
		tmap.Set(i, i)
	}

	for i := 3; i < 8; i++ {
		tmap.Del(i)
	}

	tmap.Each(func(key, value int) bool {
		fmt.Println(value)

		return true
	})

	// Output:
	// 0
	// 1
	// 2
	// 8
	// 9
}
