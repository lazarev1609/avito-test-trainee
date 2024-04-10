package models

import "time"

type Banner struct {
	ID        uint      `json:"id"`
	Data      string    `json:"data"`
	IsEnabled bool      `json:"is_enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeatureId int       `json:"feature_id"`
}
