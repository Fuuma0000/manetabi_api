package model

import "time"

type PlansCategories struct {
	PlanID     uint      `json:"plan_id"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
