package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	f, err := os.OpenFile("K:/file.mp3", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	stat, err := f.Stat() //获取文件状态
	if err != nil {
		panic(err)
	} //把文件指针指到文件末，当然你说为何不直接用 O_APPEND 模式打开，没错是可以。我这里只是试验。
	url := "http://127.0.0.1:3000/assets/37-02.mp3"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes="+strconv.FormatInt(stat.Size(), 10)+"-")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	written, err := io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}
	println("written: ", written)
}
