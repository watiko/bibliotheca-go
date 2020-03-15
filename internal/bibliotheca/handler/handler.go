package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

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

type baseStatusBody struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Commit    string `json:"commit"`
}

type failures map[string]string

type okStatusBody struct {
	baseStatusBody
}

type failedStatusBody struct {
	baseStatusBody
	Failures failures `json:"failures"`
}

func StatusHandler(commit string, db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now().Format(time.RFC3339)

		err := db.Ping()
		if err != nil {
			dbErrorMessage := err.Error()
			c.JSON(
				http.StatusInternalServerError,
				failedStatusBody{
					baseStatusBody: baseStatusBody{
						Status:    "Unavailable",
						Timestamp: now,
						Commit:    commit,
					},
					Failures: failures{
						"db": dbErrorMessage,
					},
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			okStatusBody{baseStatusBody{
				Status:    "OK",
				Timestamp: now,
				Commit:    commit,
			}},
		)
	}
}
