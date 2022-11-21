package maths

import (
	"golang.org/x/exp/constraints"
)

// Min 最小值.
func Min[Num constraints.Ordered](nums ...Num) Num {
	var min Num
	if len(nums) > 0 {
		min = nums[0]
	}

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
