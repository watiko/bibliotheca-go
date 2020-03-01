package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase"
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
	c.Status(http.StatusNotImplemented)
}

func (b Bookshelf) getAllBooksFromBookshelf(c *gin.Context) {
	_ = c.Param("book-shelf-id")
	c.Status(http.StatusNotImplemented)
}

func (b Bookshelf) createBookForBookshelf(c *gin.Context) {
	_ = c.Param("book-shelf-id")
	c.Status(http.StatusNotImplemented)
}
