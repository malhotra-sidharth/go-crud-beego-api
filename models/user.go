package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User represents a person in the system
type User struct {
	Id        int
	FirstName string    `orm:"null"`
	LastName  string    `orm:"null"`
	Email     string    `orm:"null;unique"`
	Password  string    `orm:"null"`
	RegDate   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// GetAllUsers function gets all users
func GetAllUsers() []*User {
	o := orm.NewOrm()
	var users []*User
	o.QueryTable(new(User)).All(&users)

	return users
}

// InsertOneUser inserts a single new User record
func InsertOneUser(user User) *User {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))

	// get prepared statement
	i, _ := qs.PrepareInsert()

	var u User
	// Insert
	id, err := i.Insert(&user)
	if err == nil {
		// successfully inserted
		u = User{Id: int(id)}
		err := o.Read(&u)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &u
}
