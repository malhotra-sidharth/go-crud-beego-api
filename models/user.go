package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
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

	// hash password
	hash, _ := hashPassword(user.Password)
	user.Password = hash

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

// private function : unexported to hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks hashed password
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
