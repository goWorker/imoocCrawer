package routers

import (
	"github.com/astaxie/beego"
	"itembeego/controllers"
)

func init() {
	beego.Router("/detail/?:id", &controllers.MainController{})
}
