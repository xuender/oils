package ordered

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Sort 排序.
func Sort[E constraints.Ordered](slice []E) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
