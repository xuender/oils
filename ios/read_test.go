package ios_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/ios"
)

func TestReadLine(t *testing.T) {
	t.Parallel()

	count := 0

	assert.Nil(t, ios.ReadLine("../LICENSE", func(line string) error {
		count++

		return nil
	}))
	assert.Equal(t, 21, count)
}

func TestReadLine_Error(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, ios.ReadLine("../LICENSE", func(line string) error {
		return os.ErrNotExist
	}))
}
