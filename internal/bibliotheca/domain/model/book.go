package model

import (
	"time"
)

type Book struct {
	BookID      string  `json:"id"`
	BookshelfID string  `json:"bookshelfId"`
	Title       string  `json:"title"`
	BorrowedBy  *string `json:"borrowedBy"`
	Isbn        string  `json:"isbn"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
