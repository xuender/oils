package treemap_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base/treemap"
)

func TestBytesInc(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []byte{'1', '2'}, treemap.BytesInc([]byte("11")))
	assert.Equals(t, []byte{1, 0}, treemap.BytesInc([]byte{0, 0xff}))
	assert.Equals(t, []byte{0xff, 0xff, 0xff, 0xff, 0xff}, treemap.BytesInc([]byte{0xff, 0xff}))
	assert.Equals(t, []byte{0}, treemap.BytesInc(nil))
}
