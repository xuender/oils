package base_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
)

func TestErrors_Error(t *testing.T) {
	t.Parallel()
	// nolint
	assert.Equal(t, "aa, bb", base.Errors(fmt.Errorf("aa"), fmt.Errorf("bb")).Error())
	assert.Nil(t, base.Errors())
	assert.Nil(t, base.Errors(nil))
	assert.True(t, base.Errors() == nil)
}
