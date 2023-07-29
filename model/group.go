package model

import "time"

type Group struct {
	PlanID    uint      `json:"plan_id" `
	UserID    uint      `json:"user_id" `
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
}
