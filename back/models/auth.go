package models

/*
 サインアップ
*/
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
 サインイン
*/
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
