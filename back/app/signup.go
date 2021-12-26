package app

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := DbConnect()
		defer db.Close()

		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		result, err := db.Exec("INSERT INTO users(email, password) VALUES(?, ?);", email, hashed_password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	}
}
