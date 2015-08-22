package controllers

import (
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego"
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"beego-guestbook/models"
	"time"
)

// 投稿
type GreetingController struct {
	beego.Controller
}

func (this *GreetingController) Post() {
	o := orm.NewOrm()
	greet := models.Greeting{
		Name:     this.GetString("name"),
		Comment:  this.GetString("comment"),
		CreateAt: time.Now()}
	o.Insert(&greet)
	this.Ctx.Redirect(302, "/")
}
