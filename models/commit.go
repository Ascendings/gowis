package models

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/sergi/go-diff/diffmatchpatch"
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

// New - create new commit instance
func (c Commit) NewCreateCommit(pageContent, commitMessage string, page *Page, user *User) *Commit {
	return &Commit{
		CommitHash:    Commit{}.GenerateHash(pageContent + string(time.Now().UnixNano())),
		CommitDiff:    Commit{}.CreateDiff("", pageContent),
		CommitMessage: commitMessage,
		User:          user,
		Page:          page,
	}
}

// New - create new commit instance
func (c Commit) NewEditCommit(oldContent, newContent, commitMessage string, page *Page, user *User) *Commit {
	return &Commit{
		CommitHash:    Commit{}.GenerateHash(oldContent + string(time.Now().UnixNano())),
		CommitDiff:    Commit{}.CreateDiff(oldContent, newContent),
		CommitMessage: commitMessage,
		User:          user,
		Page:          page,
	}
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

// CreateDiff - create diff message between two different texts
func (c Commit) CreateDiff(text1, text2 string) string {
	// new DiffMatchPatch instance
	dmp := diffmatchpatch.New()

	// create diffs
	diffs := dmp.DiffMain(text1, text2, true)

	// generate diff text
	patch := dmp.PatchMake(diffs)
	diffText := dmp.PatchToText(patch)

	return diffText
}

// registers model with DB
func init() {
	orm.RegisterModel(new(Commit))
}
