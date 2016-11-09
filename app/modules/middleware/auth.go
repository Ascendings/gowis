package middleware

import (
	"fmt"

	"github.com/go-macaron/session"

	"gogs.ballantine.tech/gballan1/gowis/app/models"

	"gopkg.in/macaron.v1"
)

// CheckUser - makes sure the user object is set in the session if the user_id is set
func CheckUser(ctx *macaron.Context, sess session.Store) {
	// check if the user ID has been set
	if sess.Get("user_id") != nil {
		// check if we need to set the user object
		if !(sess.Get("user") != nil) {
			// setup the user model for the DB query
			user := models.User{ID: sess.Get("user_id").(int)}

			// read the user from the database
			err := models.DB.Read(&user)
			// check for errors
			if err != nil {
				// there's been a problem
				fmt.Printf("%s", err)
			} else {
				// set the user object!
				sess.Set("user", user)
				// pass the session user object to the view
				ctx.Data["user"] = user
			}
		} else {
			// pass the session user object to the view
			ctx.Data["user"] = sess.Get("user")
		}
	}
}

// Auth - user needs to be authenticated
func Auth(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	if !(sess.Get("user_id") != nil) {
		// flash a message to the user
		f.Info("You need to be logged in to do that!")
		// redirect the user
		ctx.Redirect(ctx.URLFor("auth.login"))
	}
}

// Guest - users needs to not be authenticated
func Guest(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	if sess.Get("user_id") != nil {
		// flash a message to the user
		f.Info("Hold on hoss, you don't need that.")
		// redirect the user
		ctx.Redirect(ctx.URLFor("wiki.home"))
	}
}
