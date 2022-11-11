package oss

import "runtime"

// IsWindows 是否是windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
