package tags

import (
	"math/bits"
	"sort"
	"strconv"
	"strings"

	"github.com/xuender/oils/base"
	"golang.org/x/exp/constraints"
)

// BitMap 位图.
type BitMap[T constraints.Integer] []byte

const ten = 10

// NewBitMap 新建位图.
func NewBitMap[T constraints.Integer](elems ...T) BitMap[T] {
	t := BitMap[T]{}

	return t.Add(elems...)
}

// Add 增加位图.
func (t *BitMap[T]) Add(elems ...T) BitMap[T] {
	for _, tag := range elems {
		num, bytes := Int2Byte(tag)

		if int(num) >= len(*t) {
			*t = append(*t, make([]byte, int(num+1)-len(*t))...)
		}

		(*t)[num] |= bytes
	}

	return *t
}

// AddBitMap 增加位图.
func (t *BitMap[T]) AddBitMap(elem BitMap[T]) BitMap[T] {
	if len(elem) > len(*t) {
		*t = append(*t, make([]byte, len(elem)-len(*t))...)
	}

	for i, tag := range elem {
		(*t)[i] |= tag
	}

	return *t
}

// All 全部包含.
func (t BitMap[T]) All(elems ...T) bool {
	for _, tag := range elems {
		if !t.Has(tag) {
			return false
		}
	}

	return true
}

// Any 包含任意.
func (t BitMap[T]) Any(elems ...T) bool {
	for _, tag := range elems {
		if t.Has(tag) {
			return true
		}
	}

	return false
}

// Bytes 字节码.
func (t BitMap[T]) Bytes() []byte {
	return []byte(t)
}

// Count 位图数量.
func (t BitMap[T]) Count() int {
	count := 0
	for _, b := range t {
		count += bits.OnesCount8(b)
	}

	return count
}

// Del 删除位图.
func (t BitMap[T]) Del(elems ...T) BitMap[T] {
	for _, tag := range elems {
		n, b := Int2Byte(tag)

		if int(n) < len(t) {
			t[n] &= ^b
		}
	}

	return t
}

// DelBitMap 删除位图.
func (t BitMap[T]) DelBitMap(elem BitMap[T]) BitMap[T] {
	l := len(t)
	if l > len(elem) {
		l = len(elem)
	}

	for i := 0; i < l; i++ {
		t[i] &= ^elem[i]
	}

	return t
}

// Has 包含位图.
func (t BitMap[T]) Has(tag T) bool {
	n, b := Int2Byte(tag)

	return int(n) < len(t) && (t[n]&b) > 0
}

// Join 集合连接.
func (t BitMap[T]) Join(sep string) string {
	elems := t.Slice()

	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.FormatInt(int64(elems[0]), ten)
	}

	var builder strings.Builder

	builder.WriteString(strconv.FormatInt(int64(elems[0]), ten))

	for _, i := range elems[1:] {
		builder.WriteString(sep)
		builder.WriteString(strconv.FormatInt(int64(i), ten))
	}

	return builder.String()
}

// Load 加载.
func (t *BitMap[T]) Load(bs []byte) {
	*t = bs
}

// Slice 转换切片.
func (t BitMap[T]) Slice() []T {
	ret := base.NewSlice[T]()

	for n, b := range t {
		ret.Add(Byte2Ints(T(n), b)...)
	}

	sort.Sort(ret)

	return ret
}

// String 字符串.
func (t BitMap[T]) String() string {
	return t.Join(", ")
}

// Intersection 交集.
// nolint
func Intersection[T constraints.Integer](elems ...BitMap[T]) BitMap[T] {
	switch len(elems) {
	case 0:
		return BitMap[T]{}
	case 1:
		return elems[0]
	}

	ret := BitMap[T]{}
	ret.AddBitMap(elems[0])

	for _, elem := range elems[1:] {
		if len(ret) > len(elem) {
			ret = ret[:len(elem)]
		}

		for i := range ret {
			ret[i] &= elem[i]
		}
	}

	return ret
}
