package bibliotheca

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/handler"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/handler/controller"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

type App struct {
	*types.AppContext
}

func NewApp(env string, commit string, dbURL string) *App {
	ctx := types.NewAppContext(env, commit, dbURL)
	return &App{ctx}
}

func (app *App) Router() http.Handler {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.GET("/status", handler.StatusHandler(app.Commit))

	authRequired := e.Group("/", auth.Auth())
	{
		if app.Debug {
			authRequired.GET("/ping", handler.JwtPingHandler)
		}
	}

	apiGroup := e.Group("/v1", auth.Auth())
	{
		controller.NewBookShelf(app.AppContext).Mount(apiGroup)
		controller.NewBook(app.AppContext).Mount(apiGroup)
	}

	return e
}

func (app *App) Run(port int) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}
