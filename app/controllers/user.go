package controllers

import (
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"

	"github.com/Ascendings/gowis/app/models"
)

// UserController - wiki controller
type UserController struct {
	*Controller
}

// List - list of users
func (u UserController) List(ctx *macaron.Context, sess session.Store) {
	// users array
	var users []models.User

	// get users from the DB
	qs := models.DB.QueryTable("user")
	// order the results and put them into the array
	qs.OrderBy("-created_at").All(&users)

	// add the users to the view context
	ctx.Data["users"] = users

	// set the title
	ctx.Data["title"] = "List of Users | Gowis"
	// render view
	u.Render(ctx, "users/list")
}

// View - view a user profile
func (u UserController) View(ctx *macaron.Context, f *session.Flash) {
	userID, _ := strconv.Atoi(ctx.Params("userID"))
	// User model
	user := models.User{ID: userID}

	// find the user
	err := models.DB.Read(&user, "id")
	// check for errors
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		// let the user know the user doesn't exist
		f.Info("That user doesn't exist", false)
		// redirect the user
		ctx.Redirect(ctx.URLFor("users.list"))
	} else {
		// add the user result to the view
		ctx.Data["reqUser"] = user

		// set the title
		ctx.Data["title"] = "View User | Gowis"
		// render the view
		u.Render(ctx, "users/view")
	}
}
