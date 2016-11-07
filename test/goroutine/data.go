package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type teple struct {
	Phone string
}

func main() {
	startTime := currentTimeMillis()
	str := GetFile("./file/1.txt")
	if str == "" {
		return
	}
	aa := SplitString(str)
	if str == "" {
		fmt.Println(aa)
	}
	// data := make([]*teple, len(aa))
	// for _, v := range aa {
	// 	p := &teple{}
	// 	p.Phone = GoMd5(v.Phone)
	// 	data = append(data, p)
	// }

	// if str == "" {
	// 	fmt.Println(data)
	// }

	Getnu(aa)
	endTime := currentTimeMillis()
	s := fmt.Sprintf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))
	fmt.Println(s)
}

func SplitString(str string) []*teple {
	data := make([]*teple, 0)
	strSlice := strings.Split(str, "\n")
	for _, v := range strSlice {
		p := &teple{}
		p.Phone = v
		data = append(data, p)
	}
	return data
}

//获取文件内容
func GetFile(filename string) string {
	if !Exist(filename) {
		return ""
	}

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}
	return string(f)
}

//检测文件是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Getnu(teple []*teple) {
	c4 := make(chan string, 100)
	quit := make(chan int)
	m := len(teple)
	go func() {
		for i := 0; i < m; i++ {
			<-c4
		}
		quit <- 0
	}()

	DataHandle(teple, c4, quit)
}

func DataHandle(teple []*teple, c chan string, quit chan int) {

	for _, v := range teple {
		s := GoMd5(v.Phone)
		select {
		case c <- s:
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func GoMd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	aa := fmt.Sprintf("%x", h.Sum(nil))
	return aa
}
