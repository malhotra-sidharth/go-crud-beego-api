package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	id        int
	firstName string
	lastName  string
	email     string
	password  string
	regDate   time.Time
}

func init() {
	orm.RegisterModel(new(User))
}
