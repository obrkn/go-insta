package models

import "time"

type User struct {
	BaseModel
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	FailedAttempts int        `json:"failed_attempts"`
	LockedAt       *time.Time `json:"locked_at"`
}
