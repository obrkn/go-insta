package repositories

import (
	"database/sql"

	"github.com/obrkn/twitter/models"
)

type UserRepository interface {
	CreateUser(createUser *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(createUser *models.User) error {
	_, err := ur.db.Exec("INSERT INTO users(email, password) VALUES(?, ?);", createUser.Email, createUser.Password)
	if err != nil {
		return err
	}

	return nil
}
