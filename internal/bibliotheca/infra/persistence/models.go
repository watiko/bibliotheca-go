package persistence

import "time"

type Book struct {
	BookID      uint64  `db:"book_id,omitempty"`
	BookshelfID uint64  `db:"bookshelf_id"`
	Title       string  `db:"title"`
	BorrowedBy  *string `db:"borrowed_by"`
	Isbn        string  `db:"isbn"`

	UpdatedAt time.Time `db:"updated_at,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
}

type Bookshelf struct {
	BookshelfID uint64 `db:"bookshelf_id,omitempty"`
	GroupID     uint64 `db:"group_id"`
	Name        string `db:"name"`

	UpdatedAt time.Time `db:"updated_at,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
}
