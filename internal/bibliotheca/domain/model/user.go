package model

import "time"

type UserID string

type User struct {
	UserID UserID `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
