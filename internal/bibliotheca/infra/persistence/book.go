package persistence

import (
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

var _ repository.BookRepository = &bookRepo{}

type bookRepo struct {
	*types.AppContext
}

func NewBookRepository(ctx *types.AppContext) repository.BookRepository {
	return &bookRepo{ctx}
}

func (b bookRepo) GetBookByID(userID string, bookID string) (*model.Book, error) {
	panic("implement me")
}

func (b bookRepo) UpdateBookByID(userID string, book model.Book) (*model.Book, error) {
	panic("implement me")
}

func (b bookRepo) BorrowBookByID(userID string, bookID string) (*model.Book, error) {
	panic("implement me")
}

func (b bookRepo) ReturnBookByID(userID string, bookID string) (*model.Book, error) {
	panic("implement me")
}
