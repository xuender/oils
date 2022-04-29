package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestNumber2Bytes(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []byte{0x1}, base.Number2Bytes(1))
	assert.Equals(t, []byte{0x40, 0x42, 0xf}, base.Number2Bytes(1000000))
	assert.Equals(t, []byte{0x41, 0x42, 0xf}, base.Number2Bytes(1000001))
	assert.Equals(t, []byte{}, base.Number2Bytes(0))
	assert.Equals(t, []byte{0x1f, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x9, 0x40}, base.Number2Bytes(3.14))
}

// nolint
func FuzzNumber2Bytes(f *testing.F) {
	testcases := []int{0, 1, 2, 3, 4}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig int) {
		rev := base.Number2Bytes(orig)
		assert.Equal(t, orig, base.Bytes2Number[int](rev))
	})
}

func TestBytes2Number(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, base.Bytes2Number[int]([]byte{0x1}))
	assert.Equal(t, 0, base.Bytes2Number[int]([]byte{}))
	assert.Equal(t, 3.14, base.Bytes2Number[float64]([]byte{0x1f, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x9, 0x40}))
}
