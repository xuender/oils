package oss

import (
	"os"
	"path/filepath"
)

const DefaultFileMode os.FileMode = 0o664

func Exist(paths ...string) bool {
	filename, err := filepath.Abs(filepath.Join(paths...))
	if err != nil {
		return false
	}

	_, err = os.Stat(filename)

	return !os.IsNotExist(err)
}

func IsFile(paths ...string) bool {
	if path, err := Abs(filepath.Join(paths...)); err == nil {
		if info, err := os.Stat(path); err == nil {
			return !info.IsDir()
		}
	}

	return false
}
