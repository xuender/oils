package base

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type sorts[T constraints.Ordered, A any] struct {
	keys   []T
	slices [][]A
}

func (p sorts[T, A]) Len() int                     { return len(p.keys) }
func (p sorts[T, A]) Less(indexA, indexB int) bool { return p.keys[indexA] < p.keys[indexB] }
func (p sorts[T, A]) Swap(indexA, indexB int) {
	p.keys[indexA], p.keys[indexB] = p.keys[indexB], p.keys[indexA]

	for _, slice := range p.slices {
		if size := len(slice); size <= indexA || size <= indexB {
			continue
		}

		slice[indexA], slice[indexB] = slice[indexB], slice[indexA]
	}
}

func Sorts[T constraints.Ordered, A any](keys []T, slices ...[]A) {
	data := sorts[T, A]{keys: keys, slices: slices}

	sort.Sort(data)
}
