package oss

import (
	"os"
	"path/filepath"
	"strings"
)

func Abs(path string) (string, error) {
	if path == "~" || strings.HasPrefix(path, string([]rune{'~', os.PathSeparator})) {
		home, err := os.UserHomeDir()

		return filepath.Join(home, path[1:]), err
	}

	return filepath.Abs(path)
}
