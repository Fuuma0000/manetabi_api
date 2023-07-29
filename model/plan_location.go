package model

import "time"

type PlansLocations struct {
	PlanID     uint      `json:"plan_id"`
	LocationID uint      `json:"location_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
