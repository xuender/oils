package tags_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/tags"
)

func TestBitMap_Has(t *testing.T) {
	t.Parallel()

	mytags := tags.NewBitMap[uint32]()

	for i := 0; i < 100000; i++ {
		tag := uint32(i)

		assert.False(t, mytags.Has(tag))
		mytags.Add(tag)
		assert.True(t, mytags.Has(tag))
	}

	for i := 0; i < 100000; i++ {
		tag := uint32(i)

		assert.True(t, mytags.Has(tag))
		mytags.Del(tag)
		assert.False(t, mytags.Has(tag))
	}
}

func TestBitMap_Del(t *testing.T) {
	t.Parallel()

	tag := tags.NewBitMap(1, 2, 3, 4)
	tag.Del(2, 3)

	assert.True(t, tag.Has(1))
	assert.False(t, tag.Has(2))
	assert.False(t, tag.Has(3))
	assert.True(t, tag.Has(4))
}

func TestBitMap_All(t *testing.T) {
	t.Parallel()

	ts := tags.NewBitMap(1, 2, 3, 4)

	assert.True(t, ts.All(1, 3))
	assert.False(t, ts.All(2, 6))
}

func TestBitMap_Any(t *testing.T) {
	t.Parallel()

	ts := tags.NewBitMap(1, 2, 3, 4)

	assert.True(t, ts.Any(2, 6))
	assert.False(t, ts.Any(7, 30))
}

func TestBitMap_Slice(t *testing.T) {
	t.Parallel()

	ts := tags.NewBitMap(2, 1, 4)

	ts.Add(3)
	assert.Equals(t, []int{1, 2, 3, 4}, ts.Slice())
}

func TestBitMap_String(t *testing.T) {
	t.Parallel()

	tag := tags.NewBitMap(2, 1, 4, 33)

	assert.Equal(t, "1, 2, 4, 33", tag.String())

	tag.Del(2)
	tag.Del(1)
	tag.Del(33)
	assert.Equal(t, "4", tag.String())

	tag.Del(4)
	assert.Equal(t, "", tag.String())
}

func TestBitMap_Add(t *testing.T) {
	t.Parallel()

	tag := tags.NewBitMap(2, 1, 4)

	assert.Equal(t, 1, len(tag))

	tag.Add(10)
	assert.Equal(t, 2, len(tag))
	assert.Equal(t, 2, len(tag.Add(3)))
	assert.Equal(t, 2, len(tag.Add(7, 11)))
	assert.Equal(t, 13, len(tag.Add(100)))
}

func TestBitMap_DelBitMap(t *testing.T) {
	t.Parallel()

	ts1 := tags.NewBitMap(2, 1, 4)
	ts2 := tags.NewBitMap(2, 1, 3, 99)

	ts1.DelBitMap(ts2)

	assert.Equal(t, 1, len(ts1.Slice()))
	assert.Equals(t, []int{4}, ts1.Slice())
	assert.True(t, ts1.Has(4))

	ts1 = tags.NewBitMap(2, 1, 3, 99)
	ts2 = tags.NewBitMap(2, 1, 4)

	ts1.DelBitMap(ts2)

	assert.Equal(t, 2, len(ts1.Slice()))
	assert.Equals(t, []int{3, 99}, ts1.Slice())
	assert.True(t, ts1.Has(3))
}

func TestBitMap_AddBitMap(t *testing.T) {
	t.Parallel()

	ts1 := tags.NewBitMap(2, 1, 4)
	ts2 := tags.NewBitMap(2, 1, 3, 99)

	ts1.AddBitMap(ts2)

	assert.Equal(t, 5, len(ts1.Slice()))
}

func TestBitMap_Count(t *testing.T) {
	t.Parallel()

	tag := tags.NewBitMap(2, 1, 4, 2, 4)

	assert.Equal(t, 3, tag.Count())
	assert.Equal(t, 3, len(tag.Slice()))
}

func TestIntersection(t *testing.T) {
	t.Parallel()

	ts1 := tags.NewBitMap(2, 1, 4, 7, 2, 4, 200)
	ts2 := tags.NewBitMap(2, 1, 3, 4, 2, 18, 4)
	ts3 := tags.NewBitMap(2, 1, 4, 3)
	tag := tags.Intersection(ts1, ts2, ts3)

	assert.Equal(t, 3, tag.Count())
	assert.True(t, tag.Has(2))
	assert.True(t, tag.Has(1))
	assert.True(t, tag.Has(4))

	tag = tags.Intersection(ts1)
	assert.Equals(t, tag, ts1)

	tag = tags.Intersection[int]()
	assert.Equal(t, 0, tag.Count())
}

func TestBitMap_Bytes(t *testing.T) {
	t.Parallel()

	tag := tags.NewBitMap(2, 18)
	assert.Equal(t, 3, len(tag.Bytes()))
}

func TestBitMap_Load(t *testing.T) {
	t.Parallel()

	tag1 := tags.NewBitMap(2, 18)
	tag2 := tags.NewBitMap[int]()

	tag2.Load(tag1.Bytes())
	assert.Equal(t, tag1.String(), tag2.String())
}
