package model

import "time"

type Bookshelf struct {
	BookshelfID string `json:"bookshelfId"`
	Name        string `json:"name"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
