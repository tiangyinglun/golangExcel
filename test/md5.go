package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	Md51()
	Md()
}

func Md51() {
	h := md5.New()
	io.WriteString(h, "123456")
	aa := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(aa)
}

func Md() {
	str := "123456"
	byt := []byte(str)
	s := md5.Sum(byt)
	aa := fmt.Sprintf("%x", s)
	fmt.Println(aa)
}
