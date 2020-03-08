package model

import "time"

type Bookshelf struct {
	BookshelfID string `json:"id"`
	Name        string `json:"name"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
