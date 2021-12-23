package app

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Root() {
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
