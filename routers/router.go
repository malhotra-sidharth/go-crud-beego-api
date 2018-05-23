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
		beego.NSRouter("/all", &controllers.UserController{}, "get:GetAllUsers"),
	)

	// register namespace
	beego.AddNamespace(ns)
}
