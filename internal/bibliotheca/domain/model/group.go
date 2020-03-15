package model

import "time"

type GroupID string

type Group struct {
	GroupID GroupID `json:"id"`
	Name    string  `json:"name"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
