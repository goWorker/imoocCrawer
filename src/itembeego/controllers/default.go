package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"itembeego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("id is: %s\n", id)
	c.TplName = "index.html"
	c.Data["Photos"] = models.GetPhotos()
	c.Data["Title"] = models.GetTitle()
	c.Data["Price"] = models.GetPriceString()
}
