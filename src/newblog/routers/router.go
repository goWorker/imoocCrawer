package routers

import (
	"newblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/register",&controllers.RegisterController{})

    beego.Router("/login",&controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
}

//"get:ShowRegister;post:HandleRegister")
//beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
