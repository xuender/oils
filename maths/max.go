package maths

import (
	"golang.org/x/exp/constraints"
)

// Max 最大值.
//
// Deprecated: Use github.com/samber/lo.Max instead.
func Max[Num constraints.Ordered](nums ...Num) Num {
	var max Num
	if len(nums) > 0 {
		max = nums[0]
	}

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
