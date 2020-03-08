package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/repository"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

var _ repository.BookshelfRepository = &bookshelfRepo{}

type bookshelfRepo struct {
	*types.AppContext
	db *sqlx.DB
}

func NewBookshelfRepository(ctx *types.AppContext, db *sqlx.DB) repository.BookshelfRepository {
	return &bookshelfRepo{ctx, db}
}

	panic("implement me")
func (r bookshelfRepo) GetAllBookshelvesForUser(ctx context.Context, userID string) ([]*model.Bookshelf, error) {
}

func (r bookshelfRepo) GetAllBooksFromBookshelf(ctx context.Context, bookshelfID string) ([]*model.Book, error) {
	panic("implement me")
}

func (r bookshelfRepo) CreateBookForBookshelf(ctx context.Context, bookshelfID string, book model.Book) error {
	panic("implement me")
}
