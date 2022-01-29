package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root@tcp(%s)/twitter", os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SetHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Access-Control-Allow-Origin", os.Getenv("FRONT_HOST"))

	// For preflight
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	return w
}
