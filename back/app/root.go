package app

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Root() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/signup", SignUpHandler)
	http.HandleFunc("/token", TokenHandler)
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
