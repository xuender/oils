package base

import (
	"encoding/binary"

	"golang.org/x/exp/constraints"
)

const (
	b16 = 0x1fffff
)

// JSNumber 转换成兼容JS的数值.
func JSNumber[T constraints.Integer | constraints.Float](num T) T {
	b8 := make([]byte, Eight)
	binary.BigEndian.PutUint64(b8, uint64(num))

	h := binary.BigEndian.Uint32(b8[0:Four])
	l := binary.BigEndian.Uint32(b8[Four:])
	// nolint
	return T(uint64(h&b16)*0x100000000 + uint64(l))
}
