package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root@tcp(%s)/twitter", os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
