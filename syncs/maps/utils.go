package maps

import (
	"hash/fnv"

	"github.com/xuender/oils/base"
	"golang.org/x/exp/constraints"
)

// Group 生成分组号.
func Group[K constraints.Ordered](key K, num int) int {
	switch value := (any)(key).(type) {
	case string:
		has := fnv.New32a()
		has.Write([]byte(value))

		return int(has.Sum32()) % num
	default:
		number, _ := base.ParseIntegerAny[int](value)

		if number < 0 {
			number *= -1
		}

		return number % num
	}
}
