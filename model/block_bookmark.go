package model

import "time"

type BlockBookmark struct {
	UserID    uint      `json:"user_id"`
	BlockID   uint      `json:"block_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
