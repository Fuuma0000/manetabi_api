package model

import "time"

type PlansBookmarks struct {
	UserID    uint      `json:"user_id"`
	PlanID    uint      `json:"plan_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
