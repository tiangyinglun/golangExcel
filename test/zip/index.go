package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	rc, err := zip.OpenReader("./note.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer rc.Close()
	for _, _file := range rc.File {
		f, err := _file.Open()
		if err != nil {
			fmt.Println(err)
		}
		desfile, err := os.OpenFile(_file.Name, os.O_CREATE|os.O_WRONLY, os.ModePerm)

		if err != nil {
			fmt.Println(err)

		}
		aa, err := io.CopyN(desfile, f, int64(_file.UncompressedSize64))
		if err != nil {
			fmt.Println(err)

		}
		if aa == 0 {
			break
		}
		res, err := ioutil.ReadFile(_file.Name)
		if err != nil {
			fmt.Println(err)
		}

		content := string(res)
		ss := strings.Split(content, "\r\n")
		sss, err := json.Marshal(ss)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(sss))
	}
}
