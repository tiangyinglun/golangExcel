package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/beego?charset=utf8", 30)
}
