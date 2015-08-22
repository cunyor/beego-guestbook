package routers

import (
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego"
	"beego-guestbook/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post", &controllers.GreetingController{})
}
