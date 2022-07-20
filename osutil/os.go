package osutil

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// GetPid 获取当前进程的pid
func GetPid() int {
	// FIXME
	return 0
}

// IsWindows determines whether current OS is Windows.
func IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux determines whether current OS is Linux.
func IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsDarwin determines whether current OS is Darwin.
func IsDarwin() bool {
	return "darwin" == runtime.GOOS
}

// Pwd gets the path of current working directory.
func Pwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)

	return filepath.Dir(pwd)
}
