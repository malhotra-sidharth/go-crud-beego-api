package main

import (
	"fmt"
	_ "go-crud-beego-api/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// if in develop mode
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// get database configuration from environment variables
	dbUser := os.Getenv("db_user")
	dbPwd := os.Getenv("db_pwd")
	dbName := "go_user_mgmt"
	dbString := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=utf8"

	// Register Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register default database
	orm.RegisterDataBase("default", "mysql", dbString)

	// autosync
	// db alias
	name := "default"

	// drop table and re-create
	force := true

	// print log
	verbose := true

	// error
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
