package models

type Tweet struct {
	BaseModel
	UserId  int    `json:"user_id"`
	Message string `json:"message"`
	User    User
}

type TweetResponse struct {
	BaseModel
	Message      string `json:"message"`
	UserNickname string `json:"user_nickname`
}
