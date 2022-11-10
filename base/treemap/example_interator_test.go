package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/base/treemap"
)

func ExampleTreeMap_Each() {
	llrb := treemap.New(-1, -1)

	for i := 0; i < 5; i++ {
		llrb.Set(i, i)
	}

	llrb.Each(func(key, value int) bool {
		fmt.Println(key, value)

		return true
	})

	// Output:
	// 0 0
	// 1 1
	// 2 2
	// 3 3
	// 4 4
}
