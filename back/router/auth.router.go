package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/obrkn/twitter/controllers"
)

type AuthRouter interface {
	SetAuthRouting(api *mux.Router)
}

type authRouter struct {
	ac controllers.AuthController
}

func NewAuthRouter(ac controllers.AuthController) AuthRouter {
	return &authRouter{ac}
}

func (ar *authRouter) SetAuthRouting(api *mux.Router) {
	api.HandleFunc("/token", ar.ac.Token).Methods(http.MethodGet)
	api.HandleFunc("/signup", ar.ac.SignUp).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/signin", ar.ac.SignIn).Methods(http.MethodPost, http.MethodOptions)
}
