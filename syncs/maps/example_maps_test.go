package maps_test

import (
	"fmt"
	"strconv"

	"github.com/xuender/oils/syncs/maps"
)

func Example() {
	syncMap := maps.NewString(-1)

	for i := 0; i <= 100; i++ {
		syncMap.Set(strconv.Itoa(i), i)
	}

	list := []int{}

	syncMap.Prefix(func(key string, value int) bool {
		list = append(list, value)

		return true
	}, "1")

	fmt.Println(list)

	// Output:
	// [1 10 100 11 12 13 14 15 16 17 18 19]
}
