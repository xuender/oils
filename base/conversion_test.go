package base_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

func TestConversion(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	writer := &io.PipeWriter{}

	assert.Nil(base.Conversion[io.Reader](writer, false))
	assert.Panics(func() {
		base.Conversion[io.Reader](writer, true)
	})
	assert.NotNil(base.Conversion[io.Writer](writer, true))
}
