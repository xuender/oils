package base_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
)

// nolint
func TestRecover(t *testing.T) {
	t.Parallel()

	err := func() (err error) {
		defer base.Recover(func(e error) {
			err = e
		})
		// nolint
		panic(fmt.Errorf("xx"))
	}()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "xx")

	err = func() (err error) {
		defer base.Recover(func(e error) {
			err = e
		})
		panic("xx")
	}()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "xx")
}
