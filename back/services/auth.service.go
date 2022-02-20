package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/obrkn/twitter/models"
	"github.com/obrkn/twitter/repositories"
	"github.com/obrkn/twitter/utils/logic"
	"github.com/obrkn/twitter/utils/validation"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type AuthService interface {
	Token(w http.ResponseWriter, r *http.Request) error
	SignUp(w http.ResponseWriter, r *http.Request) (models.User, error)
	SignIn(w http.ResponseWriter, r *http.Request) (models.User, error)
	Preflight(w http.ResponseWriter, r *http.Request) error
	SendAuthResponse(w http.ResponseWriter, r *http.Request, user *models.User, conde int)
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
func (as *authService) SignIn(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	signInRequestParam := models.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// バリデーション
	if err := as.av.SignInValidate(signInRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		as.rl.SendResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return models.User{}, err
	}

	// ユーザー認証
	var user models.User
	// emailに紐づくユーザーをチェック
	if err := as.ur.GetUserByEmail(&user, r.FormValue("email")); err != nil {
		return models.User{}, err
	}

	// アカウントロック
	if user.LockedAt.Add(30 * time.Minute).After(time.Now()) {
		errMessage := "アカウントがロックされています。"
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return models.User{}, fmt.Errorf("Error: %s", "Forbidden - Locked account")
	}

	// パスワード照合
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		// 10回以上ならアカウントロック
		user.FailedAttempts++
		if user.FailedAttempts > 10 {
			*user.LockedAt = time.Now()
		}
		as.ur.UpdateUser(&user)

		errMessage := "パスワードが間違っています。"
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return models.User{}, err
	}

	user.FailedAttempts = 0
	as.ur.UpdateUser(&user)

	return user, nil
}

/*
	会員登録処理
*/
func (as *authService) SignUp(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	signUpRequestParam := models.SignUpRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// バリデーション
	if err := as.av.SignUpValidate(signUpRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		as.rl.SendResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return models.User{}, err
	}

	// 同じメールアドレスのユーザーがいないか検証
	// emailに紐づくユーザーをチェック
	if isExists, err := as.ur.ExistsUserByEmail(signUpRequestParam.Email); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return models.User{}, err
	} else if isExists {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"), http.StatusUnauthorized)
		return models.User{}, errors.Errorf("「%w」のユーザーは既に登録されています。", signUpRequestParam.Email)
	}

	// レコード作成
	var createUser models.User
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequestParam.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	createUser.Email = signUpRequestParam.Email
	createUser.Password = string(hashPassword)

	err = as.ur.CreateUser(&createUser)
	if err != nil {
		return models.User{}, err
	}

	return createUser, nil
}

/*
	プリフライト処理
*/
func (as *authService) Preflight(w http.ResponseWriter, r *http.Request) error {
	as.rl.SendResponse(w, []byte("success - preflight"), http.StatusOK)

	return nil
}

/*
	ログインAPI・会員登録APIのレスポンス送信処理
*/
func (as *authService) SendAuthResponse(w http.ResponseWriter, r *http.Request, user *models.User, code int) {

	session, err := store.Get(r, "_session")
	if err != nil {
		log.Fatal(err)
	}
	session.Options.HttpOnly = true
	session.Options.MaxAge = 60 * 60 * 1

	session.Values["Login"] = true
	session.Values["UserId"] = user.Id
	session.Values["ExpiredAt"] = time.Now().Add(time.Hour).Format("2006-01-02 15:04:05 +0900 JST")

	err = session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}

	as.rl.SendResponse(w, []byte(user.Email), code)
}
