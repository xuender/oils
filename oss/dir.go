package oss

import (
	"os"
	"path/filepath"
)

const DefaultDirFileMod os.FileMode = 0o771

func IsDir(paths ...string) bool {
	if path, err := Abs(filepath.Join(paths...)); err == nil {
		if info, err := os.Stat(path); err == nil {
			return info.IsDir()
		}
	}

	return false
}

func MkdirParent(path string) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, DefaultDirFileMod)
	}
}
