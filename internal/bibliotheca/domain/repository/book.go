package repository

import (
	"context"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
)

// TODO: precondition check
type BookRepository interface {
	NextID() model.BookID
	GetBookByID(ctx context.Context, bookID string) (*model.Book, error)
	UpdateBook(ctx context.Context, book model.Book) error
	BorrowBook(ctx context.Context, userID string, bookID string) error
	ReturnBook(ctx context.Context, userID string, bookID string) error
}
