package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/Fuuma0000/manetabi_api/model"
)

type IUserInfrastructer interface {
	CreateUser(user *model.User) error
	Login(user *model.User, email string) error
	CheckDuplicateEmail(email string) (bool, error)
	GetUserByEmail(user *model.UserResponse, email string) error
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

func (ui *userInfrastructer) Login(user *model.User, email string) error {
	q := `SELECT * FROM users WHERE email = ? LIMIT 1`
	row := ui.db.QueryRow(q, email)

	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ProfileImagePath, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("ユーザーが見つかりませんでした")
		}
		return err
	}

	return nil
}

// メールアドレスの重複チェック
func (ui *userInfrastructer) CheckDuplicateEmail(email string) (bool, error) {
	q := `SELECT COUNT(*) FROM users WHERE email = ?`
	row := ui.db.QueryRow(q, email)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (ui *userInfrastructer) GetUserByEmail(user *model.UserResponse, email string) error {
	q := `SELECT user_id, user_name, email, profile_image_path, created_at, updated_at FROM users WHERE email = ? LIMIT 1`
	row := ui.db.QueryRow(q, email)

	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.ProfileImagePath, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("ユーザーが見つかりませんでした")
		}
		return err
	}

	return nil
}
