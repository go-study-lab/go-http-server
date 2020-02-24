package handler

import (
	"fmt"
	"net/http"
)

func DisplayFormDataHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	for key, values := range r.Form {
		fmt.Fprintf(w, "Form field %q, Values %q\n", key, values)

		fmt.Fprintf(w, "Form field %q, Value %q\n", key, r.FormValue(key))
	}
}
