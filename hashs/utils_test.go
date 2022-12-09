package hashs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/hashs"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	value := 1

	assert.Equal(t, []byte{3, 4, 0, 2}, hashs.Encode(1))
	assert.Equal(t, []byte{3, 4, 0, 0x62}, hashs.Encode('1'))
	assert.Equal(t, []byte{0x31, 0x32, 0x33}, hashs.Encode("123"))
	assert.Equal(t, []byte{}, hashs.Encode(nil))
	assert.Equal(t, []byte{3, 4, 0, 2}, hashs.Encode(&value))
}

func TestDecode(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []byte{1}, hashs.Decode[[]byte](hashs.Encode([]byte{1})))
	assert.Equal(t, "123", hashs.Decode[string](hashs.Encode("123")))
	assert.Equal(t, 123, hashs.Decode[int](hashs.Encode(123)))

	type tmp struct {
		Name string
	}

	tmp1 := &tmp{Name: "ff"}
	assert.Equal(t, "ff", hashs.Decode[tmp](hashs.Encode(tmp1)).Name)
}
