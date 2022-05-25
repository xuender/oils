package base

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type sorts[T constraints.Ordered, A any] struct {
	keys   []T
	slices [][]A
}

func (p sorts[T, A]) Len() int           { return len(p.keys) }
func (p sorts[T, A]) Less(i, j int) bool { return p.keys[i] < p.keys[j] }
func (p sorts[T, A]) Swap(i, j int) {
	p.keys[i], p.keys[j] = p.keys[j], p.keys[i]

	for _, slice := range p.slices {
		if size := len(slice); size <= i || size <= j {
			continue
		}

		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Sorts[T constraints.Ordered, A any](keys []T, slices ...[]A) {
	data := sorts[T, A]{keys: keys, slices: slices}

	sort.Sort(data)
}
