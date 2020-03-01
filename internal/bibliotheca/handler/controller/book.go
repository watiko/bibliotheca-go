package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase"
)

type Book struct {
	*types.AppContext
	usecase usecase.BookUsecase
}

func NewBook(ctx *types.AppContext, usecase usecase.BookUsecase) *Book {
	return &Book{ctx, usecase}
}

func (b Book) Mount(router gin.IRouter) {
	router.GET("/books/:book-id", b.getBookByID)
	router.PUT("/books/:book-id", b.updateBookByID)
	router.PATCH("/books/:book-id/borrow", b.borrowBookByID)
	router.PATCH("/books/:book-id/return", b.returnBookByID)
}

func (b Book) getBookByID(c *gin.Context) {
	_ = c.Param("book-id")
	c.Status(http.StatusNotImplemented)
}

func (b Book) updateBookByID(c *gin.Context) {
	_ = c.Param("book-id")
	c.Status(http.StatusNotImplemented)
}

func (b Book) borrowBookByID(c *gin.Context) {
	_ = c.Param("book-id")
	c.Status(http.StatusNotImplemented)
}

func (b Book) returnBookByID(c *gin.Context) {
	_ = c.Param("book-id")
	c.Status(http.StatusNotImplemented)
}
