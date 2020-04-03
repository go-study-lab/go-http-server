package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func WriteJsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	p := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		//... handle error
	}
}
