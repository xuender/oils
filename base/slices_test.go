package base_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

func TestAppend(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{2, 3, 4}, base.Append(3, []int{1, 2}, 3, 4))
	assert.Equal(t, []int{4}, base.Append(1, []int{1, 2}, 3, 4))
	assert.Equal(t, []int{3, 4}, base.Append(2, []int{0, 1, 2}, 3, 4))
	assert.Equal(t, []int{}, base.Append(0, []int{1, 2}, 3, 4))
	assert.Equal(t, []int{2}, base.Append(1, []int{1, 2}))
	assert.Equal(t, []int{1, 2, 3, 4}, base.Append(-1, []int{1, 2}, 3, 4))
}

func FuzzAppend(f *testing.F) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testcases := []int{-1, 0, 1, 3, 4, 5}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig int) {
		t.Helper()

		rev := base.Append(orig, slice, slice...)

		if orig < 0 {
			assert.Equal(t, 20, len(rev))
			assert.Equal(t, 0, rev[0])
			assert.Equal(t, 9, rev[19])
		}

		if orig == 0 {
			assert.Equal(t, 0, len(rev))
		}

		if orig > 0 {
			if orig >= 20 {
				assert.Equal(t, 20, len(rev))
				assert.Equal(t, 0, rev[0])
				assert.Equal(t, 9, rev[19])
			} else {
				assert.Equal(t, orig, len(rev))
			}

			assert.Equal(t, 9, rev[len(rev)-1])
		}
	})
}

func TestIndex(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 1}
	assert.Equal(t, 1, base.Index(slice, []int{2, 1}))
}

func BenchmarkDel(b *testing.B) {
	array := []rune(strings.Repeat("ender", 100))

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		base.Del(array, 'd', 'e')
	}
}

func TestReplace(t *testing.T) {
	t.Parallel()

	old := []int{1, 2, 3}
	slice := base.Replace(old, []int{2, 3}, []int{7, 8}, 1)

	assert.NotEqual(t, old, slice)
}

func BenchmarkAppend(b *testing.B) {
	array := []rune(strings.Repeat("ender", 100))

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		slice := append([]rune{}, array...)

		slice[0] = '1'
	}
}

func BenchmarkAppendMake(b *testing.B) {
	array := []rune(strings.Repeat("ender", 100))

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		slice := make([]rune, len(array))
		copy(slice, array)

		slice[0] = '1'
	}
}
