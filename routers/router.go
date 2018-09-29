package routers

import (
	"installer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.SingleController{}, "get:TestShell")
	beego.Router("/testssh", &controllers.SingleController{}, "get:TestSSH")
}
