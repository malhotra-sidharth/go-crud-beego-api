// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-crud-beego-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// hello-world route
	beego.Router("/", &controllers.HelloController{}, "get:HelloWorld")

	// init namespace
	ns := beego.NewNamespace("/api/v1",

		// get all users
		beego.NSRouter("/user/all", &controllers.UserController{}, "get:GetAllUsers"),

		// add new user
		beego.NSRouter("/user", &controllers.UserController{}, "post:AddNewUser"),

		// update an existing user
		beego.NSRouter("/user", &controllers.UserController{}, "put:UpdateUser"),

		beego.NSNamespace("/user",
			// delete a user
			beego.NSRouter("/:id", &controllers.UserController{}, "delete:DeleteUser"),

			// get a user with id
			beego.NSRouter("/:id", &controllers.UserController{}, "get:GetUserById"),
		),
	)

	// register namespace
	beego.AddNamespace(ns)
}
