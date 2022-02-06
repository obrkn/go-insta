package services

import (
	"net/http"

	"github.com/obrkn/twitter/models"
	"github.com/obrkn/twitter/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(w http.ResponseWriter, r *http.Request) error
}

type authService struct {
	ur repositories.UserRepository
}

func NewAuthService(ur repositories.UserRepository) AuthService {
	return &authService{ur}
}

/*
サインアップ
*/
func (as *authService) SignUp(w http.ResponseWriter, r *http.Request) error {

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
