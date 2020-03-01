package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

type Bookshelf struct {
	*types.AppContext
}

func NewBookShelf(ctx *types.AppContext) *Bookshelf {
	return &Bookshelf{ctx}
}

// GET, POST /v1/book-shelves/{book-shelf-id}/books
func (b Bookshelf) Mount(router gin.IRouter) {
	router.GET("/book-shelves", b.getAllBookShelf)
	router.GET("/book-shelves/:book-shelf-id/books", b.getAllBooksFromBookShelf)
	router.POST("/book-shelves/:book-shelf-id/books", b.createBookForBookShelf)
}

func (b Bookshelf) getAllBookShelf(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func (b Bookshelf) getAllBooksFromBookShelf(c *gin.Context) {
	_ = c.Param("book-shelf-id")
	c.Status(http.StatusNotImplemented)
}

func (b Bookshelf) createBookForBookShelf(c *gin.Context) {
	_ = c.Param("book-shelf-id")
	c.Status(http.StatusNotImplemented)
}
