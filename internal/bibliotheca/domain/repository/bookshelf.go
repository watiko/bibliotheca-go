package repository

import "github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"

type BookshelfRepository interface {
	GetAllBookshelves(userID string) ([]*model.Bookshelf, error)
	GetAllBooksFromBookshelf(userID string, bookshelfID uint64) ([]*model.Bookshelf, error)
	CreateBookForBookshelf(userID string, book model.Book) (*model.Book, error)
}
