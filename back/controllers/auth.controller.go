package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/obrkn/twitter/services"
)

type AuthController interface {
	Token(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	as services.AuthService
}

func NewAuthController(as services.AuthService) AuthController {
	return &authController{as}
}

/*
 CSRFトークン発行
*/
func (ac *authController) Token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.Write([]byte("Success - Token created"))
}

/*
 サインアップ
*/
func (ac *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	// サインアップ
	err := ac.as.SignUp(w, r)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("byebyebye"))
}

/*
 サインイン
*/
func (ac *authController) SignIn(w http.ResponseWriter, r *http.Request) {

}
