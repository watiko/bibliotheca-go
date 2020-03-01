package bibliotheca

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/handler"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
)

type App struct {
	Debug  bool
	Commit string
}

func NewApp(env string, commit string, dbURL string) *App {
	return &App{
		Debug:  env == "dev",
		Commit: commit,
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
	e.GET("/status", handler.StatusHandler(app.Commit))

	authRequired := e.Group("/")
	authRequired.Use(auth.Auth())
	{
		if app.Debug {
			authRequired.GET("/ping", handler.JwtPingHandler)
		}
	}

	return e
}
