package services

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/obrkn/twitter/models"
	"github.com/obrkn/twitter/repositories"
	"github.com/obrkn/twitter/utils/logic"
	"github.com/obrkn/twitter/utils/validation"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Token(w http.ResponseWriter, r *http.Request) error
	SignUp(w http.ResponseWriter, r *http.Request) error
	SignIn(w http.ResponseWriter, r *http.Request) error
}

type authService struct {
	ur repositories.UserRepository
	rl logic.ResponseLogic
	av validation.AuthValidation
}

func NewAuthService(ur repositories.UserRepository, rl logic.ResponseLogic, av validation.AuthValidation) AuthService {
	return &authService{ur, rl, av}
}

/*
	トークン発行
*/
func (as *authService) Token(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	as.rl.SendResponse(w, []byte("Success - Token created"), http.StatusOK)
	return nil
}

/*
 ログイン処理
*/
func (as *authService) SignIn(w http.ResponseWriter, r *http.Request) error {
	// RequestのBodyデータを取得
	signInRequestParam := models.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// バリデーション
	if err := as.av.SignInValidate(signInRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		as.rl.SendResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return err
	}

	// ユーザー認証
	var user models.User
	// emailに紐づくユーザーをチェック
	if err := as.ur.GetUserByEmail(&user, r.FormValue("email")); err != nil {
		return err
	}

	// パスワード照合
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		errMessage := "パスワードが間違っています。"
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return err
	}

	return nil
}

/*
	会員登録処理
*/
func (as *authService) SignUp(w http.ResponseWriter, r *http.Request) error {
	// RequestのBodyデータを取得
	signUpRequestParam := models.SignUpRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// バリデーション
	if err := as.av.SignUpValidate(signUpRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		as.rl.SendResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return err
	}

	// 同じメールアドレスのユーザーがいないか検証
	// emailに紐づくユーザーをチェック
	if isExists, err := as.ur.ExistsUserByEmail(signUpRequestParam.Email); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return err
	} else if isExists {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"), http.StatusUnauthorized)
		return errors.Errorf("「%w」のユーザーは既に登録されています。", signUpRequestParam.Email)
	}

	// レコード作成
	var createUser models.User
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequestParam.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	createUser.Email = signUpRequestParam.Email
	createUser.Password = string(hashPassword)

	err = as.ur.CreateUser(&createUser)
	if err != nil {
		return err
	}

	return nil
}
