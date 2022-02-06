package app

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.Write([]byte("sdfajsdflasdfj"))
	// if r.URL.Path != "/" {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprint(w, "404 Not Found")
	// 	return
	// }
	// db := DbConnect()
	// defer db.Close()
	// w.Header().Set("X-CSRF-Token", csrf.Token(r))

	// rows, err := db.Query("SELECT id, email FROM users;")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var u User
	// for rows.Next() {
	// 	err = rows.Scan(&u.Id, &u.Email)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("ID: %v, Email: %s\n", u.Id, u.Email)
	// }

	// data1 := "hello, world"
	// outputJson, err := json.Marshal(&data1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Fprint(w, string(outputJson))
}
