package auth

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func IsLoggedIn(r *http.Request) bool {
	c := appengine.NewContext(r)
	return user.Current(c) != nil
}

func IsAdminLoggedIn(r *http.Request) bool {
	return user.IsAdmin(appengine.NewContext(r))
}
