package main

import (
	//"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type info struct {
	Str string `json:"str"`
}

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	mp := make(map[string]string)
	data := readFile("a.txt")
	for _, v := range data {
		mp[v.Str] = v.Str
	}

	data1 := readFile("b.txt")
	mak := make([]string, len(data1))
	i := 0
	for _, vs := range data1 {
		if t, v := mp[vs.Str]; v {
			if t == "" {
				continue
			}
			mak[i] = t
		} else {
			continue
		}

		i++

	}
	fileName := "tt.txt"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(f, strings.Join(mak, "\n"))

	fmt.Println(time.Now().UnixNano())

}

func readFile(path string) []*info {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	aa := make([]*info, 0)
	ff := strings.Split(string(data), "\n")
	for _, v := range ff {
		m := &info{}
		m.Str = v
		aa = append(aa, m)
	}
	return aa
}
