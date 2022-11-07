package treemap_test

import (
	"fmt"

	"github.com/xuender/oils/syncs/treemap"
)

func ExampleTreeMap() {
	tmap := treemap.New(0)

	for i := 0; i < 10; i++ {
		tmap.Set(treemap.Bytes(i), i)
	}

	for i := 7; i < 8; i++ {
		tmap.Del(treemap.Bytes(i))
	}

	for i := 0; i < 10; i++ {
		value, has := tmap.Get(treemap.Bytes(i))
		if has && value != i {
			fmt.Println(i, value)
		}
	}

	fmt.Println(tmap.Len())
	// output:
	// 9
}
