package controllers

import (
	"github.com/astaxie/beego"
)

type CameraController struct {
	beego.Controller
}

func (c *CameraController) Get() {
	c.TplName = "camera.html"
}
