package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/base/treemap"
)

func Example() {
	tmap := treemap.New(0)

	for i := 0; i < 10; i++ {
		tmap.Set([]byte{byte(i)}, i)
	}

	for i := 3; i < 8; i++ {
		tmap.Del([]byte{byte(i)})
	}

	tmap.Each(func(key []byte, value int) bool {
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
