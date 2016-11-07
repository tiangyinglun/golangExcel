package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type PhpcmsAdminRole struct {
	Roleid      int    `json:"roleid"`
	Rolename    string `json:"rolename"`
	Description string `json:"description"`
	Listorder   int    `json:"listorder"`
	Disabled    int    `json:"disabled"`
}

func main() {
	l := Query()
	js, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(js)
}

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phpcms?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func Query() []*PhpcmsAdminRole {
	rows, err := db.Query("select * from phpcms_admin_role")
	if err != nil {
		fmt.Println(err)
	}
	l := make([]*PhpcmsAdminRole, 0)
	for rows.Next() {
		m := &PhpcmsAdminRole{}
		err = rows.Scan(&m.Roleid, &m.Rolename, &m.Description, &m.Listorder, &m.Disabled)
		if err != nil {
			fmt.Println(err)
		}
		l = append(l, m)
	}
	return l
}
