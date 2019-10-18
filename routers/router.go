package routers

import (
	"test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pic2txt", &controllers.Pic2txtController{})
	beego.Router("/pic4trans", &controllers.PictranslatController{})
}
