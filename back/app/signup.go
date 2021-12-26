package app

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
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
	fmt.Println("it's connected!!")
}
