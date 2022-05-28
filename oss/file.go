package oss

import (
	"os"
	"path/filepath"
)

func Exist(files ...string) bool {
	filename, err := filepath.Abs(filepath.Join(files...))
	if err != nil {
		return false
	}

	_, err = os.Stat(filename)

	return !os.IsNotExist(err)
}
