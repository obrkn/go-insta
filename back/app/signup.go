package app

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConnect()
	defer db.Close()

	email := r.FormValue("email")
	password := r.FormValue("password")
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO users(email, password) VALUES(?, ?);", email, hashed_password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println("it's connected!!")
}
