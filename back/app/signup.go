package app

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// if r.Method == "POST" {
		hash, err := bcrypt.GenerateFromPassword([]byte("パスワード"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		// Byteで返されるので文字列に変換して表示
		fmt.Println(string(hash))
	}
}
