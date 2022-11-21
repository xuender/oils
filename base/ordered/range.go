package ordered

import (
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/maths"
)

// Range 创建一个包含从 start 到 end，但不包含 end 本身范围数字的数组.
//
// 1个参数时，end 是参数1，2个参数以上，start 是参数1，end 是参数2，step 是参数3.
//
// 如果 end 是负数，而 start 和 step 没有指定，那么 step 为 -1，start 为0.
//
// 如果 end 小于 start ，会创建一个空数组，除非指定了 step.
func Range(nums ...int) []int {
	start, end, step := 0, 0, 1

	switch len(nums) {
	case 0:
		return nil
	case 1:
		end = nums[0]

		if end < 0 {
			step = -1
		}
	case base.Two:
		start = nums[0]
		end = nums[1]
	default:
		start = nums[0]
		end = nums[1]
		step = nums[2]
	}

	length := maths.Max((end-start)/step, 0)

	if length == 0 {
		return nil
	}

	res := make([]int, length)
	index := 0

	for length > 0 {
		res[index] = start

		start += step
		index++
		length--
	}

	return res
}
