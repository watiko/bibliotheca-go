package persistence

import (
	"database/sql"
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	UserID ulid.ULID `db:"user_id"`
	Name   string    `db:"name"`
	Email  string    `db:"email"`

	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type Group struct {
	GroupID ulid.ULID `db:"group_id"`
	Name    string    `db:"name"`

	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type UserGroupMembership struct {
	UserID  ulid.ULID `db:"user_id"`
	GroupID ulid.ULID `db:"group_id"`
}

type Book struct {
	BookID      ulid.ULID      `db:"book_id"`
	BookshelfID ulid.ULID      `db:"bookshelf_id"`
	Title       string         `db:"title"`
	BorrowedBy  sql.NullString `db:"borrowed_by"`
	Isbn        string         `db:"isbn"`

	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type Bookshelf struct {
	BookshelfID ulid.ULID `db:"bookshelf_id"`
	GroupID     ulid.ULID `db:"group_id"`
	Name        string    `db:"name"`

	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}
