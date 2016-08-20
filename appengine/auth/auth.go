package auth

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func IsLoggedIn(r *http.Request) bool {
	c := appengine.NewContext(r)
	u := user.Current(c)
	return u != nil
}

func IsAdminLoggedIn(r *http.Request) bool {
	c := appengine.NewContext(r)
	return user.IsAdmin(c)
}
