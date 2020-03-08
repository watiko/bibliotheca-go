package persistence

import (
	"context"
	"fmt"

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

func (r bookshelfRepo) GetAllBookshelvesForUser(ctx context.Context, userID string) ([]*model.Bookshelf, error) {
	sql := `select bookshelves.*
	  from (select * from user_group_memberships where user_id = $1) ugm
	  inner join groups on ugm.group_id = groups.group_id
	  inner join bookshelves on groups.group_id = bookshelves.group_id`

	var dbBookshelves []*Bookshelf
	err := r.db.Select(&dbBookshelves, sql, userID)
	if err != nil {
		return nil, err
	}

	bookshelves := make([]*model.Bookshelf, len(dbBookshelves))
	for i, bookshelf := range dbBookshelves {
		bookshelves[i] = &model.Bookshelf{
			BookshelfID: fmt.Sprintf("%d", bookshelf.BookshelfID),
			Name:        bookshelf.Name,
			UpdatedAt:   bookshelf.UpdatedAt,
			CreatedAt:   bookshelf.CreatedAt,
		}
	}

	return bookshelves, nil
}

func (r bookshelfRepo) GetAllBooksFromBookshelf(ctx context.Context, bookshelfID string) ([]*model.Book, error) {
	panic("implement me")
}

func (r bookshelfRepo) CreateBookForBookshelf(ctx context.Context, bookshelfID string, book model.Book) error {
	panic("implement me")
}
