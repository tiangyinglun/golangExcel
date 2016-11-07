package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Suffix = ".xlsx"
)

//定义Excel返回类型
type DataContent struct {
	Data [][]string
}

//定义返回类型
type CallBack struct {
	RBack map[string]interface{}
}

// //处理数据
func HandleData(types int32, paramMap map[string]string) (ret string, err error) {
	var dataRet CallBack
	data := make(map[string]interface{})
	var isJson bool
	if types == 1 {
		path, ok := paramMap["path"]
		value, oks := paramMap["type"]
		//如果没则参数错
		if !ok || !oks {
			data["status"] = ParamsError
			data["message"] = Message[ParamsError]
			data["detail"] = ""
			dataRet.RBack = data
		}
		if value == "json" {
			isJson = true
		} else if value == "path" {
			isJson = false
		} else {
			data["status"] = ParamsError
			data["message"] = Message[ParamsError]
			data["detail"] = ""
			dataRet.RBack = data
		}

		dataRet, err = CallBackData(path, isJson)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		data["status"] = ParamsError
		data["message"] = Message[ParamsError]
		data["detail"] = ""
		dataRet.RBack = data
	}

	retByte, err := json.Marshal(dataRet.RBack)
	if err != nil {
		fmt.Println(err)
	}
	ret = string(retByte)
	return
}

//返回
func CallBackData(path string, isJson bool) (ret CallBack, err error) {
	data := make(map[string]interface{})
	//如果后缀名错误
	if !checkExtension(path, Suffix) {
		data["status"] = SuffixError
		data["message"] = Message[SuffixError]
		data["detail"] = ""
	}
	//如果文件不存在
	if !Exist(path) {
		data["status"] = NoExist
		data["message"] = Message[NoExist]
		data["detail"] = ""
	}
	//读取Excel
	retData, err := slove(path, &DataContent{})
	if err != nil {
		fmt.Println(err)
	}
	//判断是否返回json
	if isJson {
		data["status"] = Success
		data["message"] = Message[Success]
		data["detail"] = retData
	} else {
		strName := strconv.FormatInt(time.Now().UnixNano(), 10) + RandNum(100)
		fileName := ReadValue("createFile", "path") + strName + ".txt"
		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		w := bufio.NewWriter(f)
		for _, v := range retData.Data {
			//写入文件
			w.WriteString(strings.Join(v, "\t") + "\n")
		}
		data["status"] = Success
		data["message"] = Message[Success]
		data["detail"] = fileName
	}
	ret.RBack = data
	return
}

//解析xlsx
func slove(path string, m *DataContent) (ret *DataContent, err error) {
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
		//定义slice 类型interface
		data := make([][]string, len(sheet.Rows))
		for k, row := range sheet.Rows {
			arr := make([]string, len(row.Cells))
			for s, cell := range row.Cells {
				str, e := cell.String()
				if e != nil { //如果不等于nil 恐慌
					panic(e)
				}
				arr[s] = str
			}
			data[k] = arr
		}
		m.Data = data
	}
	return m, err
}

//验证后缀
func checkExtension(path, Suffix string) bool {
	return strings.HasSuffix(path, Suffix)
}

//检测文件是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//生成数据数
func RandNum(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Intn(num))
}
