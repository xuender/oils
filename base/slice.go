package base

// Slice 切片.
type Slice[T any] []T

// NewSlice 新建切片.
func NewSlice[T any](elems ...T) Slice[T] {
	newSlice := Slice[T]{}
	newSlice.Add(elems...)

	return newSlice
}

// Add 增加.
func (p *Slice[T]) Add(elems ...T) { *p = append(*p, elems...) }

// Clean 清空.
func (p *Slice[T]) Clean() { *p = []T{} }

// Push 顶部压入元素.
func (p *Slice[T]) Push(elems ...T) { *p = append(elems, *p...) }

// String 转换成字符串.
func (p Slice[T]) String() string { return "Slice[" + Join(p, " ") + "]" }
