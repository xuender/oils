package oss

import (
	"os/exec"
)

// Exec 执行命令.
func Exec(name string, args ...string) error {
	return exec.Command(name, args...).Run()
}

// ExecOut 获取执行结果.
func ExecOut(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	return string(output), err
}
