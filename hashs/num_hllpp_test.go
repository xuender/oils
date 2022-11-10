package hashs_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/hashs"
)

// nolint
func FuzzNumHLLPP(f *testing.F) {
	testcases := [...]int{0, 1, 2, 3, 4}
	for _, tc := range testcases {
		f.Add(tc)
	}

	var count uint64

	hll := hashs.NewNumHLLPP[int]()

	f.Fuzz(func(t *testing.T, orig int) {
		hll.Add(orig)
		count++
		assert.Equal(t, count, hll.Count())
	})
}

func TestNumHLLPP_Marshal(t *testing.T) {
	t.Parallel()

	hll := hashs.NewNumHLLPP[int]()
	hll.Add(3)
	hll.Add(1)
	assert.Equal(t, 2, hll.Count())

	newHll, err := hashs.Unmarshal[int](hll.Marshal())
	assert.Nil(t, err)
	assert.Equal(t, 2, newHll.Count())
}

func TestUnMarshal(t *testing.T) {
	t.Parallel()

	_, err := hashs.Unmarshal[int]([]byte{})
	assert.NotNil(t, err)
}
