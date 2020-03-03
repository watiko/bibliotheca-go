package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
)

type ErrorDetail struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

func NewSingleErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Errors: []ErrorDetail{
		{
			Message: message,
		},
	}}
}

func MustGetUser(c *gin.Context) *auth.User {
	user, exists := auth.GetUser(c)
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewSingleErrorResponse("server bug: cannot get user context"))
		return nil
	}
	return user
}
