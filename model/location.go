package model

import "time"

type Location struct {
	LocationID   uint      `json:"location_id" gorm:"primaryKey, autoIncrement"`
	LocationName string    `json:"location_name" gorm:"size:50;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
