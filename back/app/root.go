package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Root() {
	err := godotenv.Load(fmt.Sprintf("envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Foo := os.Getenv("FOO")
	fmt.Println(Foo)
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/signup", SignUpHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/token", TokenHandler)
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
