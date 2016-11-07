package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type box struct {
}

func main() {
	data, err := ioutil.ReadFile("fk.txt")
	if err != nil {
		fmt.Println(err)
	}
	filedata := string(data)

	//var Mymap map[string]interface{}
	filename := "fankui.txt"

	str := strings.Split(filedata, "\n")

	for _, v := range str {
		bbox := strings.Split(v, "\t")
		mtr := ""
		for k, tk := range bbox {

			if k == 0 {
				mtr += tk + "\t"
			}
			if tk == "1" {
				mtr += "0" + "\t"
			} else if tk == "0" {
				mtr += "1"
				break
			}
		}
		mtr += "\n"
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(f, mtr)

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
