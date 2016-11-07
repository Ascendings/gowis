package models

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// Commit - user group model
type Commit struct {
	ID            int `orm:"pk;auto;column(id)"`
	CommitHash    string
	CommitDiff    string    `orm:"type(text)"`
	CommitMessage string    `orm:"type(text)"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)"`

	// commit belongs to user
	User *User `orm:"rel(fk)"`
	// commit belongs to page
	Page *Page `orm:"rel(fk)"`
}

// String - string representation of page
func (c *Commit) String() string {
	return fmt.Sprintf("%s", c.CommitHash)
}

// GenerateHash - generate SHA1 hash of diff content
func (c Commit) GenerateHash(content string) string {
	// create byte array of the content
	contentBytes := []byte(content)
	// create new hasher instance
	hasher := sha1.New()

	// write byte array to hasher
	hasher.Write(contentBytes)

	// encode byte array to hex string
	contentHash := hex.EncodeToString(hasher.Sum(nil))

	// return hash
	return contentHash
}

// registers model with DB
func init() {
	orm.RegisterModel(new(Commit))
}
