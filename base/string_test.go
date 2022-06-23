package base_test

import (
	"testing"
	"unicode"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestSplit(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"a", "b", "c"}, base.Split("a,b-c", ',', '-', '.'))
	assert.Equals(t, []string{"a,b-c"}, base.Split("a,b-c"))
	assert.Equals(t, []string{"aa", "Bb", "URL", "Ok"}, base.Split("aaBbURLOk", base.SepInitialisms))
}

func TestSplitFunc(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"a", "b"}, base.SplitFunc("a  b", unicode.IsSpace))
	assert.Equals(t, []string{"a"}, base.SplitFunc("a   ", unicode.IsSpace))
	assert.Equals(t, []string{"a"}, base.SplitFunc(" a   ", unicode.IsSpace))
	assert.Equals(t, []string{"a"}, base.SplitFunc(" a", unicode.IsSpace))
	assert.Equals(t, []string{"a"}, base.SplitFunc("  a", unicode.IsSpace))
	assert.Equals(t, []string{"a"}, base.SplitFunc("  a ", unicode.IsSpace))
	assert.Equals(t, []string{"ab"}, base.SplitFunc("  ab ", unicode.IsSpace))
	assert.Equals(t, []string{"ab"}, base.SplitFunc("ab", unicode.IsSpace))
}
