package repository

import "github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"

type BookRepository interface {
	GetBookByID(userID string, bookID string) (*model.Book, error)
	UpdateBookByID(userID string, book model.Book) (*model.Book, error)
	BorrowBookByID(userID string, bookID string) (*model.Book, error)
	ReturnBookByID(userID string, bookID string) (*model.Book, error)
}
