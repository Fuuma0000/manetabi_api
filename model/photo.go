package model

import "time"

type Photo struct {
	ID        uint      `json:"id" `
	BlockID   uint      `json:"block_id" `
	PhotoPath string    `json:"photo_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
