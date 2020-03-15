package usecase

import (
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/transaction"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase/input"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase/output"
)

var _ BookshelfUsecase = &bookshelfInteractor{}

type BookshelfUsecase interface {
	GetAll(data input.BookShelvesGet) (*output.BookshelvesGet, error)
	GetBooks(data input.BooksGetFromBookshelf) (*output.BooksGetFromBookshelf, error)
	CreateBook(data input.BookCreateForBookshelf) (*output.BookCreateForBookshelf, error)
}

type bookshelfInteractor struct {
	*types.AppContext
	bookshelfRepo repository.BookshelfRepository
	bookRepo      repository.BookRepository
	transaction.Transactioner
}

func NewBookshelfInteractor(appContext *types.AppContext, bookshelfRepo repository.BookshelfRepository, bookRepo repository.BookRepository, txer transaction.Transactioner) BookshelfUsecase {
	return &bookshelfInteractor{AppContext: appContext, bookshelfRepo: bookshelfRepo, bookRepo: bookRepo, Transactioner: txer}
}

func (b bookshelfInteractor) GetAll(data input.BookShelvesGet) (*output.BookshelvesGet, error) {
	panic("implement me")
}

func (b bookshelfInteractor) GetBooks(data input.BooksGetFromBookshelf) (*output.BooksGetFromBookshelf, error) {
	panic("implement me")
}

func (b bookshelfInteractor) CreateBook(data input.BookCreateForBookshelf) (*output.BookCreateForBookshelf, error) {
	// TODO: check user is belonging to groups which owns bookshelf

	book := model.NewBook(b.bookRepo.NextID(), model.BookshelfID(data.BookshelfID), data.Title, data.Isbn)
	if err := b.bookshelfRepo.CreateBookForBookshelf(b.AppContext, book); err != nil {
		return nil, err
	}

	return &output.BookCreateForBookshelf{Book: &book}, nil
}
