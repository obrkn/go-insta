package validation

import (
	"github.com/obrkn/twitter/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuthValidation interface {
	SignInValidate(signInRequest models.SignInRequest) error
	SignUpValidate(signUpRequest models.SignUpRequest) error
}

type authValidation struct{}

func NewAuthValidation() AuthValidation {
	return &authValidation{}
}

/*
	ログインパラメータのバリデーション
*/
func (av *authValidation) SignInValidate(signInRequest models.SignInRequest) error {
	return validation.ValidateStruct(&signInRequest,
		validation.Field(
			&signInRequest.Email,
			validation.Required.Error("メールアドレス入力は必須です。"),
			validation.RuneLength(5, 40).Error("メールアドレスは 5〜40 文字です。"),
			is.Email.Error("メールアドレスの形式が間違っています。"),
		),
		validation.Field(
			&signInRequest.Password,
			validation.Required.Error("パスワード入力は必須です。"),
			validation.RuneLength(6, 20).Error("パスワードは6文字以上、20文字以内で入力してください。"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
		),
	)
}

/*
	会員登録パラメータのバリデーション
*/
func (av *authValidation) SignUpValidate(signUpRequest models.SignUpRequest) error {
	return validation.ValidateStruct(&signUpRequest,
		validation.Field(
			&signUpRequest.Email,
			validation.Required.Error("メールアドレス入力は必須です。"),
			validation.RuneLength(5, 40).Error("メールアドレスは 5〜40 文字です。"),
			is.Email.Error("メールアドレスの形式が間違っています。"),
		),
		validation.Field(
			&signUpRequest.Password,
			validation.Required.Error("パスワード入力は必須です。"),
			validation.RuneLength(6, 20).Error("パスワードは6文字以上、20文字以内で入力してください。"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
		),
	)
}
