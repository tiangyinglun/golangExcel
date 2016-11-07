package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var m []filename

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println(root)

	tt := getFileList(root)
	getlist(tt)
}

func getlist(m []*filename) {
	mp := make(map[int]interface{})
	if len(m) == 0 {
		return
	}
	for k, v := range m {

		mp[k] = v.FileName
	}

	fmt.Println(mp)
}

type files struct {
	File string
	list []filename
}

type filename struct {
	FileName string
}

func getFileList(path string) (m []*filename) {
	m = make([]*filename, 0)
	//递归取所有目录
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		aa := &filename{}
		aa.FileName = path
		m = append(m, aa)
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return m
}
