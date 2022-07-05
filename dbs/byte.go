package dbs

import "golang.org/x/exp/constraints"

const eight = uint32(8)

// Int2Byte 整数转换字节码.
func Int2Byte[T constraints.Integer](i T) (T, byte) {
	return i / T(eight), 1 << (i % T(eight))
}

// Byte2Ints 字节码转换整数.
// nolint
func Byte2Ints[T constraints.Integer](n T, b byte) []T {
	base := uint32(n) * eight
	ret := []T{}

	for i := uint32(0); i < eight; i++ {
		if b&(1<<i) > 0 {
			ret = append(ret, T(base+i))
		}
	}

	return ret
}
