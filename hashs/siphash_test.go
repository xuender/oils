package hashs_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/hashs"
)

func ExampleSipHash128() {
	fmt.Println(hashs.SipHash128([]byte("123")))

	// Output:
	// 8693645449139915215 11618447955228391416
}

func ExampleSipHash64() {
	fmt.Println(hashs.SipHash64([]byte("123")))

	// Output:
	// 9379172312344772015
}

func ExampleSipHashNumber() {
	fmt.Println(hashs.SipHashNumber([]byte("123")))
	// output:
	// 2677888159399343
}

func ExampleSipHashHex() {
	fmt.Println(hashs.SipHashHex([]byte("123")))
	// output:
	// 822983866c7d3daf
}

func TestSipHash128(t *testing.T) {
	t.Parallel()

	h1, h2 := hashs.SipHash128([]byte("123"))

	assert.Equal(t, uint64(0x78a608799f2d09cf), h1)
	assert.Equal(t, uint64(0xa13d04dbacc077f8), h2)
}

func TestSipHash64(t *testing.T) {
	t.Parallel()

	h := hashs.SipHash64([]byte("123"))

	assert.Equal(t, uint64(0x822983866c7d3daf), h)
}

func FuzzSipHashNumber(f *testing.F) {
	f.Add([]byte("test"))

	f.Fuzz(func(t *testing.T, data []byte) {
		t.Helper()

		hash := hashs.SipHashNumber(data)

		assert.LessOrEqual(t, hash, uint64(1<<53))
		assert.GreaterOrEqual(t, hash, uint64(100))
	})
}

func TestSipHash32(t *testing.T) {
	t.Parallel()

	h := hashs.SipHash32([]byte("123"))

	assert.Equal(t, uint32(0x82298386), h)
}
