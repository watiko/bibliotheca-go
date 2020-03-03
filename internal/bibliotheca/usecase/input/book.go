package input

type BookGet struct {
	UserID string
	BookID string
}

type BookUpdate struct {
	UserID string
	BookID string
	Isbn   string
	Title  string
}

type BookBorrow struct {
	UserID string
	BookID string
}

type BookReturn struct {
	UserID string
	BookID string
}
