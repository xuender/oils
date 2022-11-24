package base_test

import (
	"io"
	"testing"

	"github.com/xuender/oils/base"
)

func TestConversion(t *testing.T) {
	t.Parallel()

	writer := &io.PipeWriter{}

	base.Conversion[io.Reader](writer, false)

	// assert.Panics(t, func() {
	// 	base.Conversion[io.Reader](writer, true)
	// })

	base.Conversion[io.Writer](writer, true)
}
