package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Post struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root@tcp(db)/twitter")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, email FROM users;")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	var u User
	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v, Email: %s\n", u.Id, u.Email)
	}

	data1 := "hello, world"
	outputJson, err := json.Marshal(&data1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(outputJson))
}
