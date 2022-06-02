package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestSplit(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"a", "b", "c"}, base.Split("a,b-c", ',', '-', '.'))
	assert.Equals(t, []string{"a,b-c"}, base.Split("a,b-c"))
	assert.Equals(t, []string{"aa", "Bb", "URL", "Ok"}, base.Split("aaBbURLOk", base.SepInitialisms))
}
