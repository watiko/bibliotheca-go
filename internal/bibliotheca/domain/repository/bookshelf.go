package repository

import (
	"context"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
)

// TODO: precondition check
type BookshelfRepository interface {
	GetAllBookshelvesForUser(ctx context.Context, userID string) ([]*model.Bookshelf, error)
	GetAllBooksFromBookshelf(ctx context.Context, bookshelfID string) ([]*model.Book, error)
	CreateBookForBookshelf(ctx context.Context, book model.Book) error
}
