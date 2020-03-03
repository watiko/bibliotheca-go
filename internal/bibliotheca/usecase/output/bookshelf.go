package output

import "github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"

type BookshelvesGet struct {
	Bookshelves []*model.Bookshelf
}

type BooksGetFromBookshelf struct {
	Books []*model.Book
}

type BookCreateForBookshelf struct {
	Book *model.Book
}
