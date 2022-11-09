package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/base/treemap"
)

func ExampleTreeMap_Each() {
	llrb := treemap.New(-1)

	for i := 0; i < 5; i++ {
		llrb.Set([]byte{byte(i)}, i)
	}

	llrb.Each(func(b []byte, i int) bool {
		fmt.Println(b, i)

		return true
	})

	// Output:
	// [0] 0
	// [1] 1
	// [2] 2
	// [3] 3
	// [4] 4
}
