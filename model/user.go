package model

import "time"

// User ユーザーモデル
type User struct {
	ID           uint      `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
