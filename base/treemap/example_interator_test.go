package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/base/treemap"
)

func ExampleTreeMap_Each() {
	tmap := treemap.New(-1)

	for i := 0; i < 5; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	tmap.Each(func(b []byte, i int) bool {
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
