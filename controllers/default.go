package controllers

//import (
//	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego"
//	"html/template"
//	"log"
//	"net/http"
//)

import (
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego"
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"beego-guestbook/models"
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
