package controllers

import "github.com/astaxie/beego"

type HelloController struct {
	beego.Controller
}

func (h *HelloController) HelloWorld() {
	h.Data["json"] = map[string]string{"string": "Hello World"}
	h.ServeJSON()
}
