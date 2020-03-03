package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase/input"
)

type Bookshelf struct {
	*types.AppContext
	usecase usecase.BookshelfUsecase
}

func NewBookshelf(ctx *types.AppContext, usecase usecase.BookshelfUsecase) *Bookshelf {
	return &Bookshelf{ctx, usecase}
}

// GET, POST /v1/book-shelves/{book-shelf-id}/books
func (b Bookshelf) Mount(router gin.IRouter) {
	router.GET("/book-shelves", b.getAllBookshelf)
	router.GET("/book-shelves/:book-shelf-id/books", b.getAllBooksFromBookshelf)
	router.POST("/book-shelves/:book-shelf-id/books", b.createBookForBookshelf)
}

func (b Bookshelf) getAllBookshelf(c *gin.Context) {
	user := MustGetUser(c)
	if user == nil {
		return
	}

	out, err := b.usecase.GetAll(input.BookShelvesGet{
		UserID: user.Email,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to get all bookshelves"))
		return
	}

	c.JSON(http.StatusOK, out.Bookshelves)
}

func (b Bookshelf) getAllBooksFromBookshelf(c *gin.Context) {
	user := MustGetUser(c)
	if user == nil {
		return
	}

	bookshelfID := c.Param("book-shelf-id")
	out, err := b.usecase.GetBooks(input.BooksGetFromBookshelf{
		UserID:      user.Email,
		BookshelfID: bookshelfID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to get all books from the bookshelf"))
		return
	}

	c.JSON(http.StatusOK, out.Books)
}

func (b Bookshelf) createBookForBookshelf(c *gin.Context) {
	user := MustGetUser(c)
	if user == nil {
		return
	}

	bookshelfID := c.Param("book-shelf-id")
	var bookData BookData
	err := c.ShouldBind(&bookData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewSingleErrorResponse("invalid json body"))
		return
	}

	out, err := b.usecase.CreateBook(input.BookCreateForBookshelf{
		UserID:      user.Email,
		BookshelfID: bookshelfID,
		Isbn:        bookData.Isbn,
		Title:       bookData.Title,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to create new book for the bookshelf"))
		return
	}

	c.JSON(http.StatusOK, out.Book)
}
