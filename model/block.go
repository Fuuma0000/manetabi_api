package model

import "time"

type Block struct {
	BlockID   uint      `json:"block_id"`
	PlanID    uint      `json:"plan_id"`
	BlockName string    `json:"block_name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Memo      string    `json:"memo"`
	Cost      uint      `json:"cost"`
	Address   string    `json:"address"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
