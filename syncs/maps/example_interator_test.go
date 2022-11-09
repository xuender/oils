package maps_test

import (
	"fmt"

	"github.com/xuender/oils/syncs/maps"
)

func ExampleMaps_Each() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 10; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.Each(func(key, value int) bool {
		list = append(list, value)

		return true
	})

	fmt.Println(list)

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleMaps_Range() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 100; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.Range(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 9, 14)

	fmt.Println(list)

	// Output:
	// [9 10 11 12 13]
}

func ExampleMaps_GreateOrEqual() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 10; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.GreateOrEqual(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 6)

	fmt.Println(list)

	// Output:
	// [6 7 8 9]
}

func ExampleMaps_LessThan() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 10; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.LessThan(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 3)

	fmt.Println(list)

	// Output:
	// [0 1 2]
}

func ExampleMaps_EachDesc() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 10; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.EachDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	})

	fmt.Println(list)

	// Output:
	// [9 8 7 6 5 4 3 2 1 0]
}

func ExampleMaps_RangeDesc() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 100; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.RangeDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 9, 14)

	fmt.Println(list)

	// Output:
	// [13 12 11 10 9]
}

func ExampleMaps_GreateOrEqualDesc() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 10; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.GreateOrEqualDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 6)

	fmt.Println(list)

	// Output:
	// [9 8 7 6]
}

func ExampleMaps_LessThanDesc() {
	syncMap := maps.New[int](-1)

	for i := 0; i < 10; i++ {
		syncMap.Set(i, i)
	}

	list := []int{}

	syncMap.LessThanDesc(func(key, value int) bool {
		list = append(list, value)

		return true
	}, 3)

	fmt.Println(list)

	// Output:
	// [2 1 0]
}
