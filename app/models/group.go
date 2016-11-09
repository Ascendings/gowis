package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// Group - user group model
type Group struct {
	ID               int `orm:"pk;auto;column(id)"`
	GroupName        string
	GroupDescription string `orm:"type(text)"`
	CreatedBy        int
	CreatedAt        time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt        time.Time `orm:"auto_now;type(datetime)"`

	// many-to-many user relationship
	Users []*User `orm:"reverse(many)"`
}

// String - string representation of page
func (g *Group) String() string {
	return fmt.Sprintf("%s", g.GroupName)
}

// registers model with DB
func init() {
	orm.RegisterModel(new(Group))
}
