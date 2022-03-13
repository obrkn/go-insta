package repositories

import (
	"database/sql"

	"github.com/obrkn/twitter/models"
)

type UserRepository interface {
	GetUserByNickname(user *models.User, nickname string) error
	ExistsUserByNickname(nickname string) (bool, error)
	CreateUser(createUser *models.User) error
	UpdateUser(updateUser *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

/*
	nicknameに紐づくユーザーを取得
*/
func (ur *userRepository) GetUserByNickname(user *models.User, nickname string) error {
	if err := ur.db.
		QueryRow("SELECT id, nickname, password, failed_attempts, locked_at, created_at, updated_at FROM users WHERE nickname = ?", nickname).
		Scan(&user.Id, &user.Nickname, &user.Password, &user.FailedAttempts, &user.LockedAt, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}

	return nil
}

/*
	nicknameに紐づくユーザーが存在判定
*/
func (ur *userRepository) ExistsUserByNickname(nickname string) (bool, error) {
	var isExists bool
	if err := ur.db.
		QueryRow("SELECT EXISTS ( SELECT 1 FROM users WHERE nickname = ? LIMIT 1)", nickname).
		Scan(&isExists); err != nil {
		return isExists, err
	}

	return isExists, nil
}

/*
	ユーザーデータ新規登録
*/
func (ur *userRepository) CreateUser(createUser *models.User) error {
	_, err := ur.db.Exec("INSERT INTO users(nickname, password) VALUES(?, ?);", createUser.Nickname, createUser.Password)
	if err != nil {
		return err
	}

	return nil
}

/*
	ユーザーデータ更新
*/
func (ur *userRepository) UpdateUser(updateUser *models.User) error {
	_, err := ur.db.Exec("UPDATE users SET nickname=?, password=?, failed_attempts=?, locked_at=? WHERE id=?", updateUser.Nickname, updateUser.Password, updateUser.FailedAttempts, updateUser.LockedAt, updateUser.Id)
	if err != nil {
		return err
	}

	return nil
}
