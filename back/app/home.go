package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	db := DbConnect()
	defer db.Close()

	w = SetHeader(w)

	rows, err := db.Query("SELECT id, email FROM users;")
	if err != nil {
		log.Fatal(err)
	}

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
