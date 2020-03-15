package model

import "time"

type BookshelfID string

type Bookshelf struct {
	BookshelfID BookshelfID `json:"id"`
	Name        string      `json:"name"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
