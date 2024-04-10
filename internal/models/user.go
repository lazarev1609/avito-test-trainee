package models

import "time"

type User struct {
	Id             uint      `json:"id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
