package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/Ascendings/gravago"

	"golang.org/x/crypto/bcrypt"
)

// User - wiki user model
type User struct {
	ID           int       `orm:"pk;auto;column(id)"`
	Email        string    `orm:"unique;size(150)"`
	Username     string    `orm:"unique"`
	Password     string    `orm:"size(160)"`
	PasswordSalt string    `orm:"size(12)"`
	FirstName    string    `orm:"null"`
	LastName     string    `orm:"null"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`

	// many-to-many group relationship
	Groups []*Group `orm:"rel(m2m);rel_table(users_groups)"`

	// user owns many commits
	Commits []*Commit `orm:"reverse(many)"`

	// user creates many pages
	PagesCreated []*Page `orm:"reverse(many)"`
}

// Fullname - first name + last name
func (u User) Fullname() string {
	if u.FirstName == "" || u.LastName == "" {
		// return an empty string
		return ""
	}

	// return the full name of the user
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// GetGravatarURL - returns the user's gravatar URL based on his/her email
func (u User) GetGravatarURL() string {
	grav := gravago.New()

	// change some options
	grav.SetSize(64)       // set the icon size to 80
	grav.SetIcons("retro") // use monsterid default icons
	grav.UseHTTPS(true)    // force HTTPS

	return grav.URL(u.Email)
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

	// get the string from the byte array
	passwordHash := string(hashedPassword[:])

	// return hashedPassword
	return passwordHash
}

// CheckPassword - checks the provided password against this user's hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		// return false if there was an error
		return false
	}

	// return true if the check succeeded
	return true
}

// registers model with DB
func init() {
	orm.RegisterModel(new(User))
}
