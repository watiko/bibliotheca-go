package model

import "time"

type User struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
