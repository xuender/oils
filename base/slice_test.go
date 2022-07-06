package base_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func ExampleSlice() {
	str := base.NewSlice("c", "a")

	sort.Sort(str)
	fmt.Println(str[0])
	fmt.Println(str[1])
	fmt.Println(str.All("a", "b"))
	fmt.Println(str.Any("a", "b"))
	fmt.Println(str.Has("b"))
	fmt.Println(str.Index("c"))
	fmt.Println(str.Index("b"))
	fmt.Println(str)

	// Output:
	// a
	// c
	// false
	// true
	// false
	// 1
	// -1
	// Slice[a c]
}

func TestSlice_Has(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("1", "2", "3")

	assert.True(t, s.Has("1"))
	assert.False(t, s.Has("4"))
}

func TestSlice_All(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("1", "2", "3")

	assert.True(t, s.All("1", "3"))
	assert.False(t, s.All("1", "4"))
}

func TestSlice_Any(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("1", "2", "3")

	assert.True(t, s.Any("1", "4"))
	assert.False(t, s.Any("6", "4"))
}

func TestSlice_Index(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("1", "2", "3")

	assert.Equal(t, 1, s.Index("2"))
	assert.Equal(t, -1, s.Index("6"))
}

func TestSlice_String(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("a", "c", "b")

	assert.Equal(t, "Slice[a c b]", s.String())
	assert.Equal(t, "Slice[]", base.NewSlice[string]().String())
	assert.Equal(t, "Slice[1]", base.NewSlice(1).String())
}

func TestSlice_Del(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice("a", "a", "c")

	slice.Del("a", "c")

	assert.Equal(t, 1, len(slice))
	assert.Equal(t, "a", slice[0])

	slice = base.NewSlice[string]()

	slice.Del("a", "c")

	assert.Equal(t, 0, len(slice))

	slice = base.NewSlice("a", "a", "c")

	slice.Del()

	assert.Equal(t, 3, len(slice))
}

func TestSlice_DelAll(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice("a", "a", "c")

	slice.DelAll("a", "c")

	assert.Equal(t, 0, len(slice))

	slice = base.NewSlice("a", "a", "c")

	slice.DelAll()

	assert.Equal(t, 3, len(slice))

	slice = base.NewSlice[string]()

	slice.DelAll("a", "c")

	assert.Equal(t, 0, len(slice))
}

func TestSlice_DelAllError(t *testing.T) {
	t.Parallel()

	s := base.NewSlice("a", "a", "c")

	s.DelAll("d")

	assert.Equal(t, 3, len(s))
}

func TestSlice_Indexs(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)

	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 0, slice.Indexs([]int{1, 2}))
	assert.Equal(t, 1, slice.Indexs([]int{2, 3}))
	assert.Equal(t, -1, slice.Indexs([]int{3, 3}))
	assert.Equal(t, 0, slice.Indexs([]int{1, 2, 3}))
	assert.Equal(t, -1, slice.Indexs([]int{1, 2, 3, 4}))
}

func TestSlice_Replace(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)

	assert.Equals(t, base.NewSlice(1, 2, 3), slice.Replace([]int{1, 2}, []int{2, 2}, 0))
	assert.Equals(t, base.NewSlice(2, 2, 3), slice.Replace([]int{1, 2}, []int{2, 2}, 1))
	assert.Equals(t, base.NewSlice(2, 2, 3), slice.Replace([]int{1, 1}, []int{2, 2}, 1))
	assert.Equals(t, base.NewSlice(1, 1, 3), slice.Replace([]int{2, 2}, []int{1, 1}, 1))

	slice2 := base.NewSlice(1, 2, 3, 1, 2)
	assert.Equals(t, base.NewSlice(1, 3, 1, 2), slice2.Replace([]int{1, 2}, []int{1}, 1))

	slice3 := base.NewSlice(1, 2, 3, 1, 2)
	assert.Equals(t, base.NewSlice(1, 3, 1), slice3.Replace([]int{1, 2}, []int{1}, 2))
}

func TestSlice_ReplaceAll(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 1, 2)

	assert.Equals(t, base.NewSlice(2, 2), slice.ReplaceAll([]int{1, 2}, []int{2}))
}

func TestSlice_Equal(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	assert.True(t, slice.Equal(base.NewSlice(1, 2, 3)))
	assert.False(t, slice.Equal(base.NewSlice(1, 2)))
	assert.False(t, slice.Equal(base.NewSlice(1, 2, 3, 4)))
	assert.False(t, slice.Equal(base.NewSlice(1, 2, 2)))
}

func TestSlice_Count(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 2)
	assert.Equal(t, 1, slice.Count(1))
	assert.Equal(t, 2, slice.Count(2))
	assert.Equal(t, 0, slice.Count(3))
}

func TestSlice_Counts(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 2, 2)
	assert.Equal(t, 4, slice.Counts([]int{}))
	assert.Equal(t, 1, slice.Counts([]int{1, 2}))
	assert.Equal(t, 2, slice.Counts([]int{2, 2}))
	assert.Equal(t, 0, slice.Counts([]int{1, 1}))
	assert.Equal(t, 0, slice.Counts([]int{1, 1, 1, 1, 1}))
}

func TestSlice_Unique(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 2, 2)
	assert.Equal(t, 3, slice.Count(2))

	slice.Unique()
	assert.Equal(t, 1, slice.Count(2))
}

func TestNewSlice(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 4)
	assert.True(t, slice.Has(4))
	assert.False(t, slice.Has(3))
}

func TestSlice_Clean(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 4)
	assert.Equal(t, 3, len(slice))
	slice.Clean()
	assert.Equal(t, 0, len(slice))
}

func TestSlice_Add(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 4)
	assert.Equal(t, 3, len(slice))
	slice.Add(4)
	assert.Equal(t, 4, len(slice))
}

func TestSlice_Clip(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 2)
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 3, cap(slice))
	slice.DelAll(2)
	assert.Equal(t, 1, len(slice))
	assert.Equal(t, 3, cap(slice))
	slice.Clip()
	assert.Equal(t, 1, len(slice))
	assert.Equal(t, 1, cap(slice))
}

func TestSlice_Delete(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 2)
	assert.Equal(t, 3, len(slice))
	slice.Delete(1, 2)
	assert.Equal(t, 2, len(slice))
}

func TestSlice_Grow(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 2)
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 3, cap(slice))
	slice.Grow(5)
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 8, cap(slice))
}

func TestSlice_Insert(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	assert.Equal(t, 3, len(slice))
	slice.Insert(1, 5, 5, 5)
	assert.Equal(t, 6, len(slice))
	assert.Equals(t, []int{1, 5, 5, 5, 2, 3}, slice)
}

func TestSlice_Push(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	slice.Push(5, 5, 5)
	assert.Equals(t, []int{5, 5, 5, 1, 2, 3}, slice)
}

func TestSlice_Clone(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	bak := slice.Clone()

	assert.Equals(t, slice, bak)
}

func TestSlice_Compare(t *testing.T) {
	t.Parallel()

	sliceA := base.NewSlice(1, 2, 3)
	sliceB := base.NewSlice(1, 2, 4)

	assert.Equal(t, 1, sliceB.Compare(sliceA))
}

func TestSlice_Contains(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	assert.True(t, slice.Contains(1))
	assert.False(t, slice.Contains(4))
}

func TestSlice_Join(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	assert.Equal(t, "1,2,3", slice.Join(","))
}

func TestSlice_Len(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	assert.Equal(t, 3, slice.Len())
}

func TestSlice_Less(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	assert.True(t, slice.Less(1, 2))
}

func TestSlice_Swap(t *testing.T) {
	t.Parallel()

	slice := base.NewSlice(1, 2, 3)
	slice.Swap(1, 2)
	assert.Equals(t, []int{1, 3, 2}, slice)
}
