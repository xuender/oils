package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleMap() {
	map1 := base.NewMap[int, int]()
	map1[1] = 2
	map1[3] = 5
	fmt.Println(len(map1))

	// Output:
	// 2
}
