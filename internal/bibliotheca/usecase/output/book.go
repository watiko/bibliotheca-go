package output

import "github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"

type BookGet struct {
	Book *model.Book
}

type BookUpdate struct {
	Book *model.Book
}

type BookBorrow struct {
	Book *model.Book
}

type BookReturn struct {
	Book *model.Book
}
