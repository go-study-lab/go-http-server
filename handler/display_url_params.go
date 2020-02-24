package handler

import (
"fmt"
"net/http"
)

func DisplayUrlParamsHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		fmt.Fprintf(w, "ParamName %q, Value %q\n", k, v)
		fmt.Fprintf(w, "ParamName %q, Get Value %q\n", k, r.URL.Query().Get(k))
	}
}
