package main

import (
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego"
	"beego-guestbook/Godeps/_workspace/src/github.com/astaxie/beego/example/chat/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.WSController{})
	beego.Run()
}
