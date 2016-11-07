package main

import (
	//"fmt"
	"github.com/alecthomas/log4go"
)

func main() {
	log4go.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter()) //输出到控制台,级别为DEBUG
	log4go.AddFilter("file", log4go.DEBUG, log4go.NewFileLogWriter("test.log", false))
	log4go.Info("the time is now :%s -- %s", "213", "sad")
	defer log4go.Close()
}
