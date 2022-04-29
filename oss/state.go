package oss

import (
	"os"
	"path/filepath"
	"strings"
)

type StateCode = int

const (
	Test StateCode = iota
	Play
	Run
	Build
)

func State() StateCode {
	return StateByFile(os.Args[0])
}

func StateByFile(file string) StateCode {
	dir := filepath.Dir(file)
	base := filepath.Base(file)

	if dir == "/tmpfs" && base == "play" {
		return Play
	}

	if strings.HasPrefix(dir, os.TempDir()) && strings.Contains(dir, "go-build") {
		if filepath.Base(dir) == "exe" {
			return Run
		}

		if filepath.Ext(base) == ".test" {
			return Test
		}
	}

	return Build
}
