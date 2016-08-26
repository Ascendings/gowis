package middleware

import (
	"fmt"

	"github.com/go-macaron/session"
	"gogs.ballantine.tech/gballan1/gowis/models"
)

// CheckUser - makes sure the user object is set in the session if the user_id is set
func CheckUser(sess session.Store) {
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
			}
		}
	}
}
