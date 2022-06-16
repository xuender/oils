package base_test

import (
	"io"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestConversionPass(t *testing.T) {
	t.Parallel()

	writer := &io.PipeWriter{}

	base.ConversionPass[io.Reader](writer)
}

func TestConversionMust(t *testing.T) {
	t.Parallel()

	writer := &io.PipeWriter{}

	assert.Panics(t, func() {
		base.ConversionMust[io.Reader](writer)
	})

	base.ConversionMust[io.Writer](writer)
}
