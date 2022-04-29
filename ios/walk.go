package ios

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func WalkDir(root string, walkDirFunc fs.WalkDirFunc, level int) error {
	rootLevel := strings.Count(root, string(os.PathSeparator))

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err := walkDirFunc(path, d, err); err != nil {
			return err
		}

		if strings.Count(path, string(os.PathSeparator)) >= level+rootLevel {
			return filepath.SkipDir
		}

		return nil
	})
}
