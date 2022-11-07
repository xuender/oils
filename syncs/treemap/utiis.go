package treemap

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"

	"github.com/xuender/oils/base"
)

// Bytes 转换 []byte.
func Bytes(elem any) []byte {
	if elem == nil {
		return []byte{}
	}

	buf := &bytes.Buffer{}
	encode := gob.NewEncoder(buf)

	base.Must(encode.Encode(elem))

	return buf.Bytes()
}

// Group 生成分组号.
func Group(key []byte, max int) int {
	has := fnv.New32a()
	has.Write(key)

	return int(has.Sum32()) % max
}

const _max = 0xff

// BytesInc 字节码加一.
func BytesInc(data []byte) []byte {
	if len(data) == 0 {
		return []byte{0}
	}

	res := append([]byte{}, data...)
	size := len(res)

	for i := size - 1; i >= 0; i-- {
		if res[i] != _max {
			res[i]++

			for j := i + 1; j < size; j++ {
				res[j] = 0
			}

			return res
		}
	}

	res = append(res, []byte{0xff, 0xff, 0xff}...)

	return res
}
