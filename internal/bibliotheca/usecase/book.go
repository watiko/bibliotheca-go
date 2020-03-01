package usecase

import (
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase/input"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase/output"
)

var _ BookUsecase = &bookInteractor{}

type BookUsecase interface {
	GetBook(data input.BookGet) (*output.BookGet, error)
	UpdateBook(data input.BookUpdate) (*output.BookUpdate, error)
	BorrowBook(data input.BookBorrow) (*output.BookBorrow, error)
	ReturnBook(data input.BookReturn) (*output.BookReturn, error)
}

type bookInteractor struct {
	*types.AppContext
	bookRepo repository.BookRepository
}

func NewBookInteractor(ctx *types.AppContext, bookRepo repository.BookRepository) BookUsecase {
	bookInteractor := bookInteractor{AppContext: ctx, bookRepo: bookRepo}
	return &bookInteractor
}

func (b *bookInteractor) GetBook(data input.BookGet) (*output.BookGet, error) {
	panic("implement me")
}

func (b *bookInteractor) UpdateBook(data input.BookUpdate) (*output.BookUpdate, error) {
	panic("implement me")
}

func (b *bookInteractor) BorrowBook(data input.BookBorrow) (*output.BookBorrow, error) {
	panic("implement me")
}

func (b *bookInteractor) ReturnBook(data input.BookReturn) (*output.BookReturn, error) {
	panic("implement me")
}
