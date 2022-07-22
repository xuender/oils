package ordered

import "golang.org/x/exp/constraints"

// Contains 包含.
func Contains[E constraints.Ordered](sortSlice []E, target E) bool {
	return Index(sortSlice, target) >= 0
}
