package base_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestSplitSpace(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"aa"}, base.SplitSpace("aa"))
	assert.Equals(t, []string{"aa"}, base.SplitSpace("aa "))
	assert.Equals(t, []string{"aa"}, base.SplitSpace("\taa "))
	assert.Equals(t, []string{"aa", "bb"}, base.SplitSpace("aa bb"))
	assert.Equals(t, []string{"aa", "bb"}, base.SplitSpace("aa  bb"))
	assert.Equals(t, []string{"aa", "bb"}, base.SplitSpace("aa\tbb"))
	assert.Equals(t, []string{"aa", "bb"}, base.SplitSpace("aa \tbb"))
	assert.Equals(t, []string{"aa", "bb", "cc"}, base.SplitSpace("aa\t \tbb\tcc"))
}
