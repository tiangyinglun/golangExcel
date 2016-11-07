package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strings"
)

type PhpcmsAdminRole struct {
	Roleid      int    `json:"roleid"`
	Rolename    string `json:"rolename"`
	Description string `json:"description"`
	Listorder   int    `json:"listorder"`
	Disabled    int    `json:"disabled"`
}

type PhpcmsCollectionNode struct {
	ParNum           int    `json:"par_num"`
	UrlEnd           string `json:"url_end"`
	AuthorRule       string `json:"author_rule"`
	Lastdate         int    `json:"lastdate"`
	TitleRule        string `json:"title_rule"`
	ContentRule      string `json:"content_rule"`
	ContentPageStart string `json:"content_page_start"`
	ContentPageEnd   string `json:"content_page_end"`
	ContentPage      int    `json:"content_page"`
	CollOrder        int    `json:"coll_order"`
	Urlpage          string `json:"urlpage"`
	PagesizeEnd      int    `json:"pagesize_end"`
	TitleHtmlRule    string `json:"title_html_rule"`
	ComeformHtmlRule string `json:"comeform_html_rule"`
	TimeHtmlRule     string `json:"time_html_rule"`
	DownAttachment   int    `json:"down_attachment"`
	Siteid           int    `json:"siteid"`
	Sourcecharset    string `json:"sourcecharset"`
	UrlStart         string `json:"url_start"`
	TimeRule         string `json:"time_rule"`
	ContentHtmlRule  string `json:"content_html_rule"`
	CustomizeConfig  string `json:"customize_config"`
	Name             string `json:"name"`
	PageBase         string `json:"page_base"`
	Sourcetype       int    `json:"sourcetype"`
	UrlContain       string `json:"url_contain"`
	ContentPageRule  int    `json:"content_page_rule"`
	ContentNextpage  string `json:"content_nextpage"`
	ComeformRule     string `json:"comeform_rule"`
	Watermark        int    `json:"watermark"`
	Nodeid           int    `json:"nodeid"`
	PagesizeStart    int    `json:"pagesize_start"`
	UrlExcept        string `json:"url_except"`
	AuthorHtmlRule   string `json:"author_html_rule"`
}

type DescStruct struct {
	Field string
	Type  string
}

func main() {
	Query()
}

func DescSruct() {
	db, err := sql.Open("mysql", "root:@/phpcms?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("desc phpcms_collection_node")
	if err != nil {
		fmt.Println(err)
	}
	l := make([]*DescStruct, 0)
	for rows.Next() {
		m := &DescStruct{}
		err = rows.Scan(&m.Field, &m.Type)
		if err != nil {
			fmt.Println(err)
		}
		l = append(l, m)
	}
	HandleData(l)
}

func HandleData(l []*DescStruct) {
	for k, v := range l {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func Query() {
	db, err := sql.Open("mysql", "root:@/phpcms?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	table := "phpcms_collection_node"
	rows, err := db.Query("desc " + table)
	if err != nil {
		fmt.Println(err)
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	l := make([]DescStruct, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		var m DescStruct
		m.Type = record["Type"]
		m.Field = record["Field"]
		l = append(l, m)

	}
	hander(l, table)
}

func hander(l []DescStruct, table string) {
	smp := make(map[string]string, len(l))
	for _, v := range l {
		if strings.Contains(v.Field, "_") {
			s := handerFamt(v.Field)
			if strings.Contains(v.Type, "int") {
				smp[s] = "int---" + v.Field
			} else {
				smp[s] = "string---" + v.Field
			}

		} else {
			st := strings.Title(v.Field)
			if strings.Contains(v.Type, "int") {
				smp[st] = "int---" + v.Field
			} else {
				smp[st] = "string---" + v.Field
			}
		}
	}
	str := " type " + handerFamt(table) + " struct{ \r\n "
	for k, v := range smp {
		tstr := strings.Split(v, "---")
		str += k + " " + tstr[0] + " `json:\"" + tstr[1] + "\"`" + "\r\n"
	}
	str += "}"
	filename := table + ".txt"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(f, str)
}

func handerFamt(s string) string {
	vb := strings.Split(s, "_")
	str := strings.Replace(strings.Title(strings.Join(vb, " ")), " ", "", -1)
	return str
}
