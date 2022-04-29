package ios_test

import (
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/ios"
)

func TestWalkDir(t *testing.T) {
	t.Parallel()

	_ = ios.WalkDir("/tmp", func(path string, d fs.DirEntry, err error) error {
		assert.LessOrEqual(t, len(filepath.SplitList(path)), 3)

		return nil
	}, 2)

	count := 0

	_ = ios.WalkDir("/tmp", func(path string, d fs.DirEntry, err error) error {
		count++

		return nil
	}, 0)

	assert.Equal(t, 1, count)
}
