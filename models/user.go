package models

import (
  "fmt"
  "time"

  "github.com/astaxie/beego/orm"
)

// User - wiki user model
type User struct {
  ID        int `orm:"pk;auto;column(id)"`
  Email     string
  Username  string `orm:"unique"`
  Password  string
  FirstName string    `orm:"null"`
  LastName  string    `orm:"null"`
  CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
  UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// Fullname - first name + last name
func (u *User) Fullname() string {
  if u.FirstName == "" || u.LastName == "" {
    // return an empty string
    return ""
  }

  // return the full name of the user
  return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// registers model with DB
func init() {
  orm.RegisterModel(new(User))
}
