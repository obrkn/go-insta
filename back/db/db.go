package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root@tcp(%s)/twitter?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
