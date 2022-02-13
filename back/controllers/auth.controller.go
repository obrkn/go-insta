package controllers

import (
	"fmt"
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
		fmt.Println(err)
		w.Write([]byte("サインアップ失敗です"))
		return
	}

	w.Write([]byte("サインアップ成功です"))
}

/*
	ログイン
*/
func (ac *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	// ログイン
	err := ac.as.SignIn(w, r)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ログイン失敗です"))
		return
	}

	w.Write([]byte("ログイン成功です。"))
}
