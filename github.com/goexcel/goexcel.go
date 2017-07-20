package goexcel

import (
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
	"os"
	"bufio"
	"strings"
	"encoding/json"
)

// //处理数据
func HandleData(types int32, paramMap map[string]string, c *CallBack) (ret string, err error) {
	if types == 1 {
		//excel 转成 json或 txt
		ret, err = HandleExcelFile(paramMap, c)
		return
	} else if types == 2 {
		//json ， txt to excel
		ret, err = HandleFileExcel(paramMap, c)
		return
	}
	c.RBack["status"] = ParamsError
	c.RBack["message"] = Message[ParamsError]
	c.RBack["detail"] = ""
	ret, err = c.RanderJson()
	return

}

/**
 生成excel
 */
func HandleFileExcel(paramMap map[string]string, c *CallBack) (ret string, err error) {

	var data [][]string
	backTypes, status := paramMap["type"]
	if !status {
		c.RBack["status"] = ParamsError
		c.RBack["message"] = Message[ParamsError]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
		return
	}
	if backTypes != "json" && backTypes != "path" {
		c.RBack["status"] = ParamsError
		c.RBack["message"] = Message[ParamsError]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
		return
	}
	if backTypes == "path" {
		path, ok := paramMap["path"]
		if !ok { //是否存在该参数
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
			return
		}
		if !Exist(path) { //验证文件是否存在
			c.RBack["status"] = NoExist
			c.RBack["message"] = Message[NoExist]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
			return
		}
		if !checkExtension(path, Suftxt) {
			c.RBack["status"] = SuffixErr
			c.RBack["message"] = Message[SuffixErr]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
			return
		}
		//处理txt文件
		dataContent, err := readCluesFileLine(path)
		if err != nil {
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err1 := c.RanderJson()
			return ret, err1
		}
		ret, err1 := CreateExcel(dataContent, c)
		return ret, err1
	} else {
		//传入json
		js, isbol := paramMap["json"]
		if !isbol {
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
			return
		}
		jsbyte := []byte(js)
		err := json.Unmarshal(jsbyte, &data)
		if err != nil {
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err1 := c.RanderJson()
			return ret, err1
		}
		ret, err1 := CreateExcel(data, c)
		return ret, err1
	}
	c.RBack["status"] = NoKnowErr
	c.RBack["message"] = Message[NoKnowErr]
	c.RBack["detail"] = ""
	ret, err = c.RanderJson()
	return
}

/**
生成excel
 */
func CreateExcel(d [][]string, c *CallBack) (ret string, err error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	for _, v := range d {
		row = sheet.AddRow()
		for _, vt := range v {
			cell = row.AddCell()
			cell.Value = vt
		}
	}
	var rootPath string
	//生成MD5名子
	strName := Md5(strconv.FormatInt(time.Now().UnixNano(), 10)+RandNum(1000)) + Sufxlsx
	rootPath = ReadValue("filepath", "filepath")
	if rootPath == "" {
		rootPath = GetEnvPath()
	}
	//文件名称
	fileName := rootPath + "/" + strName
	err = file.Save(fileName)
	if err != nil {
		c.RBack["status"] = CreateFail
		c.RBack["message"] = Message[CreateFail]
		c.RBack["detail"] = ""
		ret, err1 := c.RanderJson()
		return ret, err1
	}
	c.RBack["status"] = Success
	c.RBack["message"] = Message[Success]
	c.RBack["detail"] = fileName
	ret, err = c.RanderJson()
	return
}

func HandleExcelFile(paramMap map[string]string, c *CallBack) (ret string, err error) {
	var backtype bool
	path, ok := paramMap["path"]
	if !ok { //是否存在该参数
		c.RBack["status"] = ParamsError
		c.RBack["message"] = Message[ParamsError]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
		return
	}
	if !Exist(path) { //验证文件是否存在
		c.RBack["status"] = NoExist
		c.RBack["message"] = Message[NoExist]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
		return
	}
	if !checkExtension(path, Sufxlsx) {
		c.RBack["status"] = SuffixErr
		c.RBack["message"] = Message[SuffixErr]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
		return
	}
	backTypes, status := paramMap["type"]
	if !status {
		backtype = true
	} else {
		if backTypes == "json" {
			backtype = true
		} else if backTypes == "path" {
			backtype = false
		} else {
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
			return
		}
	}
	result, err := slove(path, &DataContent{})
	if err != nil {
		c.RBack["status"] = NoKnowErr
		c.RBack["message"] = Message[NoKnowErr]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
		return
	}
	ret, err = CallBackCreateData(backtype, result, c)
	return
}

func CallBackCreateData(backtype bool, m *DataContent, c *CallBack) (ret string, err error) {
	var rootPath string
	if backtype {
		c.RBack["status"] = Success
		c.RBack["message"] = Message[Success]
		c.RBack["detail"] = m.Data
		ret, err = c.RanderJson()
		return
	}
	//生成MD5名子
	strName := Md5(strconv.FormatInt(time.Now().UnixNano(), 10)+RandNum(1000)) + Suftxt
	rootPath = ReadValue("filepath", "filepath")
	if rootPath == "" {
		rootPath = GetEnvPath()
	}
	//文件名称
	fileName := rootPath + "/" + strName
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()
	if err != nil {
		c.RBack["status"] = CreateFail
		c.RBack["message"] = Message[CreateFail]
		c.RBack["detail"] = m.Data
		ret, err = c.RanderJson()
		return
	}
	buf := bufio.NewWriter(f)
	for _, v := range m.Data {
		buf.WriteString(strings.Join(v, "\t") + "\n")
	}
	buf.Flush()
	if err != nil {
		c.RBack["status"] = NoKnowErr
		c.RBack["message"] = Message[NoKnowErr]
		c.RBack["detail"] = m.Data
		ret, err = c.RanderJson()
		return
	}
	c.RBack["status"] = Success
	c.RBack["message"] = Message[Success]
	c.RBack["detail"] = fileName
	ret, err = c.RanderJson()
	return
}

//解析xlsx
func slove(paths string, m *DataContent) (ret *DataContent, err error) {
	//打开文件
	xlFile, err := xlsx.OpenFile(paths)
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
			arr := make([]string, len(sheet.Rows[0].Cells))
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
