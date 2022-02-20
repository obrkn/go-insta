package controllers

import (
	"net/http"

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
	// プリフライト
	if r.Method == http.MethodOptions {
		ac.as.Preflight(w, r)
		return
	}

	// トークン発行
	ac.as.Token(w, r)
}

/*
 サインアップ
*/
func (ac *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	// プリフライト
	if r.Method == http.MethodOptions {
		ac.as.Preflight(w, r)
		return
	}

	// サインアップ
	user, err := ac.as.SignUp(w, r)
	if err != nil {
		return
	}

	// レスポンス送信
	ac.as.SendAuthResponse(w, r, &user, http.StatusOK)
}

/*
	ログイン
*/
func (ac *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	// プリフライト
	if r.Method == http.MethodOptions {
		ac.as.Preflight(w, r)
		return
	}

	// ログイン
	user, err := ac.as.SignIn(w, r)
	if err != nil {
		return
	}

	// レスポンス送信
	ac.as.SendAuthResponse(w, r, &user, http.StatusOK)
}
