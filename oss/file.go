package oss

import (
	"os"
	"path/filepath"
)

const DefaultFileMode os.FileMode = 0o664

func Exist(files ...string) bool {
	filename, err := filepath.Abs(filepath.Join(files...))
	if err != nil {
		return false
	}

	_, err = os.Stat(filename)

	return !os.IsNotExist(err)
}
