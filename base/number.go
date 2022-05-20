package base

import (
	"encoding/binary"

	"golang.org/x/exp/constraints"
)

const (
	b16 = 0x1fffff
)

// JSNumber 转换成兼容JS的数值.
func JSNumber[Num constraints.Integer | constraints.Float](num Num) Num {
	b8 := make([]byte, Eight)
	binary.BigEndian.PutUint64(b8, uint64(num))

	h := binary.BigEndian.Uint32(b8[0:Four])
	l := binary.BigEndian.Uint32(b8[Four:])
	// nolint
	return Num(uint64(h&b16)*0x100000000 + uint64(l))
}

// Max 最大值.
func Max[Num constraints.Integer | constraints.Float](nums ...Num) Num {
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

// Min 最小值.
func Min[Num constraints.Integer | constraints.Float](nums ...Num) Num {
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
