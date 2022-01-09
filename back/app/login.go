package app

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	if r.Method == "POST" {
		session, _ := store.Get(r, "cookie-name")

		// Authentication goes here
		// ...

		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)
	}
	fmt.Println("it's connected!!")
}
