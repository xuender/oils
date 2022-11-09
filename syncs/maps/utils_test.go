package maps_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/maps"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	value := 1

	assert.Equals(t, []byte{3, 4, 0, 2}, maps.Encode(1))
	assert.Equals(t, []byte{3, 4, 0, 0x62}, maps.Encode('1'))
	assert.Equals(t, []byte{0x31, 0x32, 0x33}, maps.Encode("123"))
	assert.Equals(t, []byte{}, maps.Encode(nil))
	assert.Equals(t, []byte{3, 4, 0, 2}, maps.Encode(&value))
}

func TestDecode(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []byte{1}, maps.Decode[[]byte](maps.Encode([]byte{1})))
	assert.Equal(t, "123", maps.Decode[string](maps.Encode("123")))
	assert.Equal(t, 123, maps.Decode[int](maps.Encode(123)))

	type tmp struct {
		Name string
	}

	tmp1 := &tmp{Name: "ff"}
	assert.Equal(t, "ff", maps.Decode[tmp](maps.Encode(tmp1)).Name)
}
