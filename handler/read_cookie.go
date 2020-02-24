package handler

import (
	"fmt"
	"net/http"
)

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, "Cookie field %q, Value %q\n", cookie.Name, cookie.Value)
	}
}
