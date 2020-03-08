package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

var _ repository.BookRepository = &bookRepo{}

type bookRepo struct {
	*types.AppContext
	db *sqlx.DB
}

func NewBookRepository(ctx *types.AppContext, db *sqlx.DB) repository.BookRepository {
	return &bookRepo{ctx, db}
}

func (b bookRepo) GetBookByID(ctx context.Context, bookID string) (*model.Book, error) {
	panic("implement me")
}

func (b bookRepo) UpdateBook(ctx context.Context, book model.Book) error {
	panic("implement me")
}

func (b bookRepo) BorrowBook(ctx context.Context, userID string, bookID string) error {
	panic("implement me")
}

func (b bookRepo) ReturnBook(ctx context.Context, userID string, bookID string) error {
	panic("implement me")
}
