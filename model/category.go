package model

import "time"

type Category struct {
	CategoryID   uint      `json:"category_id" `
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
