package bibliotheca

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
)

type App struct {
	Debug bool
}

func jwtPingHandler(app *App) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func (app *App) Router() http.Handler {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "hello, world",
			},
		)
	})

	authRequired := e.Group("/")
	authRequired.Use(auth.Auth())
	{
		if app.Debug {
			authRequired.GET("/ping", jwtPingHandler(app))
		}
	}

	return e
}
