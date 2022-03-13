package models

/*
 サインアップ
*/
type SignUpRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

/*
 サインイン
*/
type SignInRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
