package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	verifyKey = []byte(os.Getenv("SESSION_KEY"))
	store     = sessions.NewCookieStore(verifyKey)
)

func Root() {
	err := godotenv.Load(fmt.Sprintf("envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	csrfMiddleware := csrf.Protect([]byte(verifyKey), csrf.Secure(false))
	api := r.PathPrefix("/api").Subrouter()
	api.Use(csrfMiddleware)

	api.HandleFunc("/home", HomeHandler).Methods(http.MethodGet)
	api.HandleFunc("/signup", SignUpHandler).Methods(http.MethodPost)
	api.HandleFunc("/login", LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":8080", r)
	// r.Use(mux.CORSMethodMiddleware(r))
	// if os.Getenv("GO_ENV") == "development" {
	// 	log.Fatal(http.ListenAndServe(":8080", csrf.Protect(verifyKey, csrf.Secure(false))(r)))
	// } else {
	// 	log.Fatal(http.ListenAndServe(":8080", csrf.Protect(verifyKey)(r)))
	// }
}
