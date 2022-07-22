package ordered

import (
	"github.com/xuender/oils/base"
	"golang.org/x/exp/constraints"
)

// Index 查找.
func Index[E constraints.Ordered](sortSlice []E, target E) int {
	start, end := 0, len(sortSlice)-1

	for start <= end {
		index := (start + end) / base.Two

		switch {
		case target == sortSlice[index]:
			return index
		case target > sortSlice[index]:
			start = index + 1
		default:
			end = index - 1
		}
	}

	return -1
}
