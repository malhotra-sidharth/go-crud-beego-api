package controllers

import "github.com/astaxie/beego"

// HelloController represents controller for Hello-World API
type HelloController struct {
	beego.Controller
}

func (h *HelloController) HelloWorld() {
	h.Data["json"] = map[string]string{"string": "Hello World"}
	h.ServeJSON()
}
