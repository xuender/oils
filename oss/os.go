package oss

import (
	"os"
	"os/exec"
)

// Open 打开文件.
func Open(file string) error {
	switch {
	case IsDarwin():
		return exec.Command("open", file).Start()
	case IsWindows():
		return exec.Command("cmd", "/k", "start", file).Start()
	case IsLinux():
		return exec.Command("xdg-open", file).Start()
	default:
		return ErrNotDefined
	}
}

// Show 在目录中显示.
func Show(file string) error {
	switch {
	case IsDarwin():
		return exec.Command("open", "-R", file).Start()
	case IsWindows():
		info, err := os.Stat(file)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return exec.Command("cmd", "/k", "explorer", file).Start()
		}

		return exec.Command("cmd", "/k", "explorer", "/select", file).Start()
	case IsLinux():
		return exec.Command("nautilus", file).Start()
	default:
		return ErrNotDefined
	}
}
