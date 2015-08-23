package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"../models"
)

// ---
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	o := orm.NewOrm()

	var greetings []*models.Greeting
	o.QueryTable("greeting").OrderBy("-CreateAt").All(&greetings)

	this.Data["greetings"] = greetings
	this.TplNames = "index.tpl"
}
