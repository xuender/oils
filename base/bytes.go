package base

import (
	"encoding/binary"

	"golang.org/x/exp/constraints" // nolint
)

func Number2Bytes[T constraints.Integer | constraints.Float](num T) []byte {
	data := make([]byte, Eight)
	binary.LittleEndian.PutUint64(data, uint64(num))

	for i := Eight; i > 0; i-- {
		if data[i-1] > 0 {
			return data[:i]
		}
	}

	return []byte{}
}

func Bytes2Number[T constraints.Integer | constraints.Float](data []byte) T {
	bytes := make([]byte, Eight)
	copy(bytes, data)
	// nolint
	return T(binary.LittleEndian.Uint64(bytes))
}
