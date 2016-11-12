package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
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

type HeadTitle struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Ext  int    `json:"ext"`
}

// //处理数据
func HandleData(types int32, paramMap map[string]string, c *CallBack) (ret string, err error) {

	var isJson bool
	if types == 1 {
		path, ok := paramMap["path"]
		value, oks := paramMap["type"]
		//如果没则参数错
		if !ok || !oks {
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
		}
		if value == "json" {
			isJson = true
		} else if value == "path" {
			isJson = false
		} else {
			c.RBack["status"] = ParamsError
			c.RBack["message"] = Message[ParamsError]
			c.RBack["detail"] = ""
			ret, err = c.RanderJson()
		}
		ret, err = CallBackData(path, isJson, c)
	} else if types == 2 {

	} else {
		c.RBack["status"] = ParamsError
		c.RBack["message"] = Message[ParamsError]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
	}
	return
}

//验证检测返回
func CallbackCheck(paramMap map[string]string, c *CallBack) (ret string, err error) {
	subscript, ok := paramMap["subscript"]
	mark, oks := paramMap["mark"]
	head, okc := paramMap["head"]
	path, okp := paramMap["path"]
	//验证参数
	if !ok || oks || okc || !okp {
		c.RBack["status"] = ParamsError
		c.RBack["message"] = Message[ParamsError]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
	}
	//解析传过来的参数 subscript
	subscriptByte := []byte(subscript)
	var subBox []int
	err = json.Unmarshal(subscriptByte, &subBox)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(subBox)
	//解析参数 mark
	markBox := make(map[string][]int)
	markByte := []byte(mark)
	err = json.Unmarshal(markByte, &markBox)
	if err != nil {
		fmt.Println(err)
	}
	//解析参数 head
	headByte := []byte(head)
	var headbox []HeadTitle
	err = json.Unmarshal(headByte, &headbox)
	if err != nil {
		fmt.Println(err)
	}
	//读取文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	//存放所有的错误
	//Message := make(map[string]map[string]string)
	i := 0
	for {
		i++
		// if i == 1 {
		// 	bfRd.ReadString('\n')
		// 	continue
		// }
		line, err := bfRd.ReadString('\n')
		//行数下标
		errLine := strconv.Itoa(i)
		if err != nil || io.EOF == err { //遇到任何错误立即返回，并忽略 EOF 错误信息
			break
		}
		//把字符串变成数组
		data := strings.Split(strings.TrimRight(line, "\n"), "\t")
		//存错误
		m := make(map[string]string)
		//简单验证电话号码
		phone, err := strconv.ParseInt(data[0], 10, 0)
		//验证电话号码
		if err != nil || ((phone < 13000000000 && phone < 16000000000) || (phone < 17000000000 && phone > 19000000000) && len(data[0]) != 32) {
			m["phone"] = "1"
			c.RBack[errLine] = m
		}
		//定义是否已经该结束了
		var iszero bool = false
		dataLen := len(data)
		//循环阶段
		for k, v := range subBox {
             //如果长度小于第一个下标 跳出循环
			if dataLen<v&&k==0{
				m["level"] = "1"
				c.RBack[errLine] = m
				break
			}
			//如果长多小于v 就代表阶段要结束
			if dataLen < v {
				break
			}
	       //如果最后一个 为空 跳出 并且结束了 就跳出
           if dataLen==v&&data[v]==""&&iszero {
           	    continue
           }
			//已经结束 后面不为空|| 如果 内容为空 并且没有结束 
			if ( iszero && data[v]!="") || data[v]==""&& !iszero{
				m["level"] = "1"
				c.RBack[errLine] = m
				break
			}
		
			dataV := data[v]
			//转成数字32 位
			value, err := strconv.Atoi(data[v])
			//入托不是数字 就跳出
			if err != nil {
				m["level"] = "1"
				c.RBack[errLine] = m
				break
			}
			//比较数量 和value值的大小
			lmark := len(markBox[headbox[v].Type]) < value
			if lmark {
				m["level"] = "1"
				c.RBack[errLine] = m
				break
			}
			//如果已经结束
			if iszero {
				m["level"] = "1"
				c.RBack[errLine] = m
				break
			}
			//如果 值不是0 就true
			if dataV != "0" {
				iszero = true
			}
		} 

	}

	ret, err = c.RanderJson()
	return

}

//验证 text 的格式是否正确
func CallBackData(path string, isJson bool, c *CallBack) (ret string, err error) {
	//如果后缀名错误
	if !checkExtension(path, Suffix) {
		c.RBack["status"] = SuffixError
		c.RBack["message"] = Message[SuffixError]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
	}
	//如果文件不存在
	if !Exist(path) {
		c.RBack["status"] = NoExist
		c.RBack["message"] = Message[NoExist]
		c.RBack["detail"] = ""
		ret, err = c.RanderJson()
	}
	//读取Excel
	retData, err := slove(path, &DataContent{})
	if err != nil {
		fmt.Println(err)
	}
	//判断是否返回json
	if isJson {
		c.RBack["status"] = Success
		c.RBack["message"] = Message[Success]
		c.RBack["detail"] = retData.Data
		ret, err = c.RanderJson()
	} else {
		//生成文件名
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
		w.Flush()
		c.RBack["status"] = Success
		c.RBack["message"] = Message[Success]
		c.RBack["detail"] = fileName
		ret, err = c.RanderJson()
	}
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

//check json的内容是否合法

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

//把 c 转成json
func (c *CallBack) RanderJson() (jsonStr string, err error) {
	jsonS, err := json.Marshal(c.RBack)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr = string(jsonS)
	return
}
