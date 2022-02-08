package services

import (
	"net/http"

	"github.com/obrkn/twitter/models"
	"github.com/obrkn/twitter/repositories"
	"github.com/obrkn/twitter/utils/logic"
	"github.com/obrkn/twitter/utils/validation"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
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
 ログイン処理
*/
func (as *authService) SignIn(w http.ResponseWriter, r *http.Request) error {
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

	// レコード作成
	var createUser models.User
	createUser.Email = r.FormValue("email")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	createUser.Password = string(hashPassword)

	err = as.ur.CreateUser(&createUser)
	if err != nil {
		return err
	}

	return nil
}
