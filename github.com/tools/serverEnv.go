package tools

import (
	"os"
	"runtime"
)

//获取系统类型 如windows linux
func GetOs() string {
	return runtime.GOOS
}

//获取程序路径
func GetPath() string {
	return os.Getenv("PWD")
}
