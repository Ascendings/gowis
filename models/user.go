package models

import (
  "bytes"
  "fmt"
  "time"

  "github.com/astaxie/beego/orm"

  "golang.org/x/crypto/bcrypt"
)

// User - wiki user model
type User struct {
  ID           int       `orm:"pk;auto;column(id)"`
  Email        string    `orm:"size(150)"`
  Username     string    `orm:"unique"`
  Password     string    `orm:"size(128)"`
  PasswordSalt string    `orm:"size(12)"`
  FirstName    string    `orm:"null"`
  LastName     string    `orm:"null"`
  CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
  UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
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

// HashPassword - hashes a provided password
func (u User) HashPassword(password string) string {
  // create byte array of the password
  passwordBytes := []byte(password)

  // Hashing the password with the default cost of 10
  hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
  if err != nil {
    panic(err)
  }

  // get the length of the byte array
  n := bytes.IndexByte(hashedPassword, 0)
  // get the string from the byte array
  passwordHash := string(hashedPassword[:n])

  // return hashedPassword
  return passwordHash
}

// registers model with DB
func init() {
  orm.RegisterModel(new(User))
}
