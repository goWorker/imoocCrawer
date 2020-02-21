package controllers

import (
	"beegocms/hellobeego/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	models.UpdatePage()
	m := models.GetPage()
	c.Data["Website"] = m.Website
	c.Data["Email"] = m.Email
	c.TplName = "index.tpl"

}
