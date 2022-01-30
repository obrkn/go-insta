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

	csrfMiddleware := csrf.Protect(
		[]byte(verifyKey),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"*localhost*"}),
	)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(csrfMiddleware)

	api.HandleFunc("/token", token).Methods(http.MethodGet)
	api.HandleFunc("/home", HomeHandler).Methods(http.MethodGet)
	api.HandleFunc("/home", HomeHandler).Methods(http.MethodPost)
	api.HandleFunc("/post", post).Methods(http.MethodPost)
	api.HandleFunc("/post", post).Methods(http.MethodOptions)
	api.HandleFunc("/signup", SignUpHandler).Methods(http.MethodPost)
	api.HandleFunc("/login", LoginHandler).Methods(http.MethodPost)
	// r.HandleFunc("/api/token", token).Methods(http.MethodGet)
	// r.HandleFunc("/api/post", post)
	r.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":8080", r)
	// r.Use(mux.CORSMethodMiddleware(r))
	// if os.Getenv("GO_ENV") == "development" {
	// 	log.Fatal(http.ListenAndServe(":8080", csrf.Protect(verifyKey, csrf.Secure(false))(r)))
	// } else {
	// 	log.Fatal(http.ListenAndServe(":8080", csrf.Protect(verifyKey)(r)))
	// }
}

func token(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Add("Access-Control-Expose-Headers", "X-CSRF-Token")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.Write([]byte("hello, rosk"))
}
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	// if r.Method == "POST" {
	// 	db := DbConnect()
	// 	defer db.Close()

	// 	_, err := db.Exec("INSERT INTO test(content) VALUES(?);", "OK, boomer")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Fprint(w, "ok your command")
	// }
}
