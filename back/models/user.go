package models

import "time"

type User struct {
	BaseModel
	Nickname       string     `json:"nick_name"`
	Password       string     `json:"password"`
	FailedAttempts int        `json:"failed_attempts"`
	LockedAt       *time.Time `json:"locked_at"`
}

type UserResponse struct {
	Nickname string `json:"nick_name"`
}
