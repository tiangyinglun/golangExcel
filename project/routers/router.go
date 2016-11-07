package routers

import (
	"github.com/astaxie/beego"
	"project/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.AutoRouter(&controllers.ShowController{})
}
