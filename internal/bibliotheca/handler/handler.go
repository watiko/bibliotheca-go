package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
)

func JwtPingHandler(c *gin.Context) {
	user, exists := auth.GetUser(c)

	if !exists {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"email":  user.Email,
				"header": user.Token.Header,
				"claims": user.Token.Claims,
			},
		)

	}
}

func StatusHandler(commit string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":    "OK",
				"timestamp": time.Now().Format(time.RFC3339),
				"system": gin.H{
					"commit": commit,
				},
			},
		)
	}
}
