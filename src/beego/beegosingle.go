package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RESTfulController struct {
	beego.Controller
}

type RegExpController struct {
	beego.Controller
}

func (this *RESTfulController) Get() {
	this.Ctx.WriteString("Hello world in get method")
}

func (this *RESTfulController) Post() {
	this.Ctx.WriteString("Hello world in POST method")
}

func (this *RegExpController) Get() {
	this.Ctx.WriteString(fmt.Sprintf("In regexp mode !\n"))
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(fmt.Sprintf("id is: %s\n", id))

	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString(fmt.Sprintf("splat is :%s\n", splat))

	path := this.Ctx.Input.Param(":path")
	this.Ctx.WriteString(fmt.Sprintf("path is :%s\n", path))

	ext := this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString(fmt.Sprintf("ext is :%s\n", ext))
}

func main() {
	beego.Router("/RESTFul", &RESTfulController{})
	beego.Router("/RegExp1/?:id", &RegExpController{})
	beego.Router("/RegExp2/:id([0-9]+)", &RegExpController{})
	beego.Router("/RegExp3/:id([\\w]+)", &RegExpController{})
	beego.Router("/RegExp4/abc:id([0-9]+)de", &RegExpController{})
	beego.Router("/RegExp5/*", &RegExpController{})
	beego.Router("/RegExp6/*.*", &RegExpController{})
	beego.Run("127.0.0.1:8082")
}
