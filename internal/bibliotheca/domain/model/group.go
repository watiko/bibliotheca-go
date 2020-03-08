package model

import "time"

type Group struct {
	GroupID string `json:"id"`
	Name    string `json:"name"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
