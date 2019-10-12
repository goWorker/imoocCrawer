package routers

import (
	"github.com/astaxie/beego"
	"todo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
