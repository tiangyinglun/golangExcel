package main

import (
	//"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	//	"log"
	//	"net/http"
	"os"
	"strings"
	//	"time"
)

func main() {
	// http.HandleFunc("/", handleData)         // 设置访问的路由
	// err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	tt := slove("./demo.xlsx", &dats{})
	fmt.Println(tt)

}

type dats struct {
	aa [][]string
}

// //处理数据
// func handleData(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() // 解析参数，默认是不会解析的
// 	if len(r.Form) > 0 {
// 		path := strings.Join(r.Form["path"], "")
// 		//验证后缀
// 		if checkExtension(path) {
// 			if Exist(path) { //验证是否存在
// 				fmt.Println(time.Now().Unix())
// 				str := slove(path)
// 				//输出到浏览器
// 				fmt.Fprintf(w, str)
// 				fmt.Println(time.Now().Unix())
// 			} else {
// 				err := `{"status":2,"message":"文件不存在，检查后重试"}`
// 				fmt.Fprintf(w, err)
// 			}
// 		} else {
// 			err := `{"status":1,"message":"参数地址必须以 .xlsx 为后缀的文件"}`
// 			fmt.Fprintf(w, err)
// 		}

// 	}

// }

//解析xlsx
func slove(path string, m *dats) *dats {
	//打开文件
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		panic(err)
	}
	//循环 sheet
	for sk, sheet := range xlFile.Sheets {
		if sk > 0 {
			break
		}
		c := make([][]string, len(sheet.Rows))
		for k, row := range sheet.Rows {
			b := make([]string, len(row.Cells))
			for t, cell := range row.Cells {
				str, e := cell.String()
				if e != nil { //如果不等于nil 恐慌
					panic(e)
				}
				b[t] = str
			}
			c[k] = b
		}
		m.aa = c
	}
	return m
}

//验证后缀
func checkExtension(path string) bool {
	return strings.HasSuffix(path, ".xlsx")
}

//检测文件是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
