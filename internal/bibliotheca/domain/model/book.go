package model

import (
	"time"
)

type BookID string

type Book struct {
	BookID      BookID      `json:"id"`
	BookshelfID BookshelfID `json:"bookshelfId"`
	Title       string      `json:"title"`
	BorrowedBy  *string     `json:"borrowedBy"`
	Isbn        string      `json:"isbn"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewBook(bookID BookID, bookshelfID BookshelfID, title string, isbn string) Book {
	return Book{
		BookID:      bookID,
		BookshelfID: bookshelfID,
		Title:       title,
		Isbn:        isbn,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
}
