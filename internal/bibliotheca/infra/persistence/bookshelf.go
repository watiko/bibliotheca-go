package persistence

import (
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

var _ repository.BookshelfRepository = &bookshelfRepo{}

type bookshelfRepo struct {
	*types.AppContext
}

func NewBookshelfRepository(ctx *types.AppContext) repository.BookshelfRepository {
	return &bookshelfRepo{ctx}
}

func (r bookshelfRepo) GetAllBookshelves(userID string) ([]*model.Bookshelf, error) {
	panic("implement me")
}

func (r bookshelfRepo) GetAllBooksFromBookshelf(userID string, bookshelfID uint64) ([]*model.Bookshelf, error) {
	panic("implement me")
}

func (r bookshelfRepo) CreateBookForBookshelf(userID string, book model.Book) (*model.Book, error) {
	panic("implement me")
}
