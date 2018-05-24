package controllers

import (
	"encoding/json"
	"go-crud-beego-api/models"

	"github.com/astaxie/beego"
)

// UserController represents controller for user API
type UserController struct {
	beego.Controller
}

// GetAllUsers function gets all the users
func (uc *UserController) GetAllUsers() {
	uc.Data["json"] = models.GetAllUsers()
	uc.ServeJSON()
}

// AddNewUser adds new user
func (uc *UserController) AddNewUser() {
	var u models.User
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	user := models.InsertOneUser(u)
	uc.Data["json"] = user
	uc.ServeJSON()
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser() {
	var u models.User
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	user := models.UpdateUser(u)
	uc.Data["json"] = user
	uc.ServeJSON()
}
