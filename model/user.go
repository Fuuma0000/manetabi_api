package model

import "time"

// User ユーザーモデル
type User struct {
	ID               uint      `json:"id"`
	UserName         string    `json:"user_name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	ProfileImagePath *string   `json:"profile_image_path"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID               uint      `json:"id"`
	UserName         string    `json:"user_name"`
	Email            string    `json:"email"`
	ProfileImagePath *string   `json:"profile_image_path"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
