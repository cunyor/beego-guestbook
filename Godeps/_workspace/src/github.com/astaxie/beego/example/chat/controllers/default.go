package controllers

import (
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["host"] = this.Ctx.Request.Host
	this.TplNames = "index.tpl"
}
