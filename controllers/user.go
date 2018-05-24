package controllers

import (
	"encoding/json"
	"go-crud-beego-api/models"
	"strconv"

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

// DeleteUser deletes an existing user
func (uc *UserController) DeleteUser() {
	// get id from query string and convert it to int
	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))

	// delete user
	deleted := models.DeleteUser(id)

	// generate response
	uc.Data["json"] = map[string]bool{"deleted": deleted}
	uc.ServeJSON()
}

// GetUserById gets a single user with the given id
func (uc *UserController) GetUserById() {
	// get the id from query string
	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))

	// get user
	user := models.GetUserById(id)

	// generate response
	uc.Data["json"] = user
	uc.ServeJSON()
}
