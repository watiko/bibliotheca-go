package bibliotheca

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/handler"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/handler/controller"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/infra/persistence"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/middleware/auth"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/usecase"
)

type App struct {
	*types.AppContext
	bookController      *controller.Book
	bookshelfController *controller.Bookshelf
}

func NewApp(env string, commit string, dbURL string) *App {
	ctx := types.NewAppContext(env, commit, dbURL)

	bookRepo := persistence.NewBookRepository(ctx)
	bookUsecase := usecase.NewBookInteractor(ctx, bookRepo)
	bookController := controller.NewBook(ctx, bookUsecase)

	bookshelfRepo := persistence.NewBookshelfRepository(ctx)
	bookshelfUsecase := usecase.NewBookshelfInteractor(ctx, bookshelfRepo)
	bookshelfController := controller.NewBookshelf(ctx, bookshelfUsecase)

	return &App{
		AppContext:          ctx,
		bookController:      bookController,
		bookshelfController: bookshelfController,
	}
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
		app.bookController.Mount(apiGroup)
		app.bookshelfController.Mount(apiGroup)
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
