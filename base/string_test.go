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

func TestLevenshteinDistance(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, base.LevenshteinDistance("a", "a"))
	assert.Equal(t, 1, base.LevenshteinDistance("a", ""))
	assert.Equal(t, 1, base.LevenshteinDistance("", "a"))
	assert.Equal(t, 3, base.LevenshteinDistance("horse", "ros"))
	assert.Equal(t, 5, base.LevenshteinDistance("intention", "execution"))
}

func TestWildcard(t *testing.T) {
	t.Parallel()

	assert.False(t, base.Wildcard("aa", "a"))
	assert.True(t, base.Wildcard("aa", "*"))
	assert.False(t, base.Wildcard("cb", "?a"))
	assert.True(t, base.Wildcard("adceb", "a*b"))
	assert.False(t, base.Wildcard("acdcb", "a*c?b"))
	assert.True(t, base.Wildcard("a", "*a*"))
}
