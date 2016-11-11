package tools

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//获取系统类型 如windows linux
func GetOs() string {
	return runtime.GOOS
}

//获取程序路径
func GetPath() string {
	file, _ := exec.LookPath(os.Args[0])
	return filepath.Dir(file)
}
