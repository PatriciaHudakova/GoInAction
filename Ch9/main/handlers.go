package main

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandlerFunc("/sendjsonn", SendJSON) //bind endpoint to function
}

//returns simple json document
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@email.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
