package main

import (
	"bufio"
	"fmt"
	//	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())
	strName := RandNum(1000) + strconv.Itoa(time.Now().Nanosecond()) + RandNum(1000)
	//生成文件
	jsonStr := strName + ".txt"
	f, err := os.OpenFile(jsonStr, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(f)

	for i := 0; i < 100000; i++ {
		str := strconv.Itoa(i) + "夜空中最  亮的星你是否明白我   是谁啊啊哈哈哈哈......" + "\n"
		w.WriteString(str)
	}

	fmt.Println(time.Now().UnixNano())
}

func RandNum(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Intn(num))
}
