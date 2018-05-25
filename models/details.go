package models

import "github.com/astaxie/beego/orm"

// Detail represents phone number and address
// of a user. A user can have multiple numbers and addresses.
// But only one primary addresses
type Detail struct {
	Id          int
	PhoneNumber int
	Primary     bool
	User        *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Detail))
}
