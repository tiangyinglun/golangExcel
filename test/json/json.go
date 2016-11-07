package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type info struct {
	Hello string `json:"hello"`
	Nihao string `json:"nihao"`
}

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

func main() {
	Test()
}

//把json 转化成map
func Test() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	//把字符串变成byte
	strb := []byte(jsonStr)
	jsonArr := make(map[string]interface{})
	err := json.Unmarshal(strb, &jsonArr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonArr)
	fmt.Println(jsonArr["host"])
	fmt.Println("================map 到json str=====================")
	jst, err := json.Marshal(jsonArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jst))
}

//把json转化到struct
func Test2() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	str := []byte(jsonStr)
	var conf ConfigStruct
	err := json.Unmarshal(str, &conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}

func Test3() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	str := []byte(jsonStr)
	var other Other
	err := json.Unmarshal(str, &other)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(other)
}

func Test4() {
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

}

func Jsons() {
	mp := make(map[string]interface{})
	mp["hello"] = "i want fly"
	mp["nihao"] = " i dont not know "
	js, err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", js)
}

func dejson() {
	str := `{"hello":"i want fly","nihao":" i dont not know "}`
	//先解析成byte
	sby := []byte(str)
	var infos info
	err := json.Unmarshal(sby, &infos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(infos.Hello)
}
