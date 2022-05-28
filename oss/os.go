package oss

import (
	"os"
	"os/exec"
	"runtime"
)

// Open 打开文件.
func Open(file string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", file).Start()
	case "windows":
		return exec.Command("cmd", "/k", "start", file).Start()
	case "linux":
		return exec.Command("xdg-open", file).Start()
	default:
		return ErrNotDefined
	}
}

// Show 在目录中显示.
func Show(file string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", "-R", file).Start()
	case "windows":
		info, err := os.Stat(file)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return exec.Command("cmd", "/k", "explorer", file).Start()
		}

		return exec.Command("cmd", "/k", "explorer", "/select", file).Start()
	case "linux":
		return exec.Command("nautilus", file).Start()
	default:
		return ErrNotDefined
	}
}
