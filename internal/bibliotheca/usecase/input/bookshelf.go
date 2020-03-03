package input

type BookShelvesGet struct {
	UserID string
}

type BooksGetFromBookshelf struct {
	UserID      string
	BookshelfID string
}

type BookCreateForBookshelf struct {
	UserID      string
	BookshelfID string
	Isbn        string
	Title       string
}
