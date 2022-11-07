package treemap_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/treemap"
)

func TestBytes(t *testing.T) {
	t.Parallel()

	value := 1

	assert.Equals(t, []byte{3, 4, 0, 2}, treemap.Bytes(1))
	assert.Equals(t, []byte{3, 4, 0, 0x62}, treemap.Bytes('1'))
	assert.Equals(t, []byte{6, 0xc, 0, 3, 0x31, 0x32, 0x33}, treemap.Bytes("123"))
	assert.Equals(t, []byte{}, treemap.Bytes(nil))
	assert.Equals(t, []byte{3, 4, 0, 2}, treemap.Bytes(&value))
	assert.Equals(t, []byte{
		0xc, 0xff, 0x81, 0x2, 0x1, 0x2, 0xff,
		0x82, 0x0, 0x1, 0x4, 0x0, 0x0, 0x4, 0xff, 0x82, 0x0, 0x0,
	}, treemap.Bytes([]int{}))
}

func TestBytesInc(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []byte{'1', '2'}, treemap.BytesInc([]byte("11")))
	assert.Equals(t, []byte{1, 0}, treemap.BytesInc([]byte{0, 0xff}))
	assert.Equals(t, []byte{0xff, 0xff, 0xff, 0xff, 0xff}, treemap.BytesInc([]byte{0xff, 0xff}))
	assert.Equals(t, []byte{0}, treemap.BytesInc(nil))
}
