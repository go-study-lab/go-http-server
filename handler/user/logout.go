package user

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, sessionCookieName)

	session.Values["authenticated"] = false
	session.Save(r, w)
}
