package dbs

import (
	"math/bits"

	"github.com/xuender/oils/base"
)

// Tag 标签.
//
// 标签内最多保存64个状态，数据库索引字段避免使用Tag，位运算无法使用索引加速.
type Tag uint64

// Add 增加标签.
func (p *Tag) Add(nums ...int) {
	for _, num := range nums {
		*p |= 1 << num
	}
}

// Del 删除标签.
func (p *Tag) Del(nums ...int) {
	for _, num := range nums {
		*p &= ^(1 << num)
	}
}

// Len 标签数量.
func (p Tag) Len() int {
	return bits.OnesCount64(uint64(p))
}

// Has 包含.
func (p Tag) Has(num int) bool {
	return p&(1<<num) > 0
}

// Slice 转切片.
func (p Tag) Slice() []int {
	res := []int{}

	for i := 0; i < base.SixtyFour; i++ {
		if p.Has(i) {
			res = append(res, i)
		}
	}

	return res
}

// All 全部包含.
func (p Tag) All(nums ...int) bool {
	for _, num := range nums {
		if !p.Has(num) {
			return false
		}
	}

	return true
}

// Any 包含任意个.
func (p Tag) Any(nums ...int) bool {
	for _, num := range nums {
		if p.Has(num) {
			return true
		}
	}

	return false
}
