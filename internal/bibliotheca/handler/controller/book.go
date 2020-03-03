package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase/input"
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
	user := MustGetUser(c)
	if user == nil {
		return
	}

	bookID := c.Param("book-id")

	out, err := b.usecase.GetBook(input.BookGet{
		UserID: user.Email,
		BookID: bookID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to get the book"))
		return
	}

	c.JSON(http.StatusOK, out.Book)
}

func (b Book) updateBookByID(c *gin.Context) {
	user := MustGetUser(c)
	if user == nil {
		return
	}

	bookID := c.Param("book-id")
	var bookData BookData
	err := c.ShouldBind(&bookData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewSingleErrorResponse("invalid json body"))
		return
	}

	out, err := b.usecase.UpdateBook(input.BookUpdate{
		UserID: user.Email,
		BookID: bookID,
		Isbn:   bookData.Isbn,
		Title:  bookData.Title,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to update the book"))
		return
	}

	c.JSON(http.StatusOK, out.Book)
}

func (b Book) borrowBookByID(c *gin.Context) {
	user := MustGetUser(c)
	if user == nil {
		return
	}

	bookID := c.Param("book-id")

	out, err := b.usecase.BorrowBook(input.BookBorrow{
		UserID: user.Email,
		BookID: bookID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to borrow the book"))
		return
	}

	c.JSON(http.StatusOK, out.Book)
}

func (b Book) returnBookByID(c *gin.Context) {
	user := MustGetUser(c)
	if user == nil {
		return
	}

	bookID := c.Param("book-id")

	out, err := b.usecase.ReturnBook(input.BookReturn{
		UserID: user.Email,
		BookID: bookID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("failed to return the book"))
		return
	}

	c.JSON(http.StatusNotImplemented, out.Book)
}
