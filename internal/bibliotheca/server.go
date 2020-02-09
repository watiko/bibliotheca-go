package bibliotheca

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "hello, world",
			},
		)
	})

	return e
}
