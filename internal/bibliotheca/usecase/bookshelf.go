package usecase

import (
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
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
}

func NewBookshelfInteractor(appContext *types.AppContext, bookshelfRepo repository.BookshelfRepository) BookshelfUsecase {
	return &bookshelfInteractor{AppContext: appContext, bookshelfRepo: bookshelfRepo}
}

func (b bookshelfInteractor) GetAll(data input.BookShelvesGet) (*output.BookshelvesGet, error) {
	panic("implement me")
}

func (b bookshelfInteractor) GetBooks(data input.BooksGetFromBookshelf) (*output.BooksGetFromBookshelf, error) {
	panic("implement me")
}

func (b bookshelfInteractor) CreateBook(data input.BookCreateForBookshelf) (*output.BookCreateForBookshelf, error) {
	panic("implement me")
}
