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
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONT_HOST"))
	return w
}
