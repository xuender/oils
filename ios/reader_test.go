package ios_test

import (
	"bytes"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/ios"
)

func TestNewContainsReader(t *testing.T) {
	t.Parallel()

	reader := bytes.NewBuffer([]byte("12333456"))
	creader := ios.NewContainsReader(reader, []byte("333"))
	data := make([]byte, 10)
	_, _ = creader.Read(data)

	assert.True(t, creader.Contains())
}
