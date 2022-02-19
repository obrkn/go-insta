package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type MainRouter interface {
	setupRouting() *mux.Router
	StartWebServer() error
}

type mainRouter struct {
	authR AuthRouter
}

func NewMainRouter(authR AuthRouter) MainRouter {
	return &mainRouter{authR}
}

/*
 ルーティング定義
*/
func (mainRouter *mainRouter) setupRouting() *mux.Router {
	router := mux.NewRouter()

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("SESSION_KEY")),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"*localhost*"}),
	)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(csrfMiddleware)

	mainRouter.authR.SetAuthRouting(api)
	api.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("obaobabarara"))
		rw.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		rw.Header().Add("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		rw.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		rw.Header().Add("Access-Control-Expose-Headers", "X-CSRF-Token")
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
	}).Methods(http.MethodPost, http.MethodOptions)

	return router
}

/*
 サーバー起動
*/
func (mainRouter *mainRouter) StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")

	return http.ListenAndServe(":8080", mainRouter.setupRouting())
}
