package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fileName := "test.txt"
	createFile(fileName)

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		m := strconv.Itoa(i)
		io.WriteString(f, m+": hello world \n")
	}

}

//检测文件是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//创建文件

func createFile(dir string) {
	if Exist(dir) != true {
		dstFile, err := os.Create(dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer dstFile.Close()
	}
}
