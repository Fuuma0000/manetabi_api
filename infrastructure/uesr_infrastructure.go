package infrastructure

import (
	"database/sql"

	"github.com/Fuuma0000/manetabi_api/model"
)

type IUserInfrastructer interface {
	CreateUser(user *model.User) error
}

type userInfrastructer struct {
	db *sql.DB
}

func NewUserInfrastructer(db *sql.DB) IUserInfrastructer {
	return &userInfrastructer{db}
}

func (ui *userInfrastructer) CreateUser(user *model.User) error {
	q := `INSERT INTO users (user_name, email, password, profile_image_path) VALUES (?, ?, ?, ?)`
	_, err := ui.db.Exec(q, user.UserName, user.Email, user.Password, user.ProfileImagePath)
	if err != nil {
		return err
	}
	return nil
}
