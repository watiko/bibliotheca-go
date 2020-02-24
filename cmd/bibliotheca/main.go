package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/sync/errgroup"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca"
	"github.com/watiko/bibliotheca-go/internal/pkg"
)

var commit string

type Env struct {
	pkg.DBEnv
	Port int    `default:"8080"`
	Env  string `default:"dev"`
}

func main() {
	var env Env

	err := envconfig.Process("APP", &env)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	var eg errgroup.Group

	if commit == "" {
		commit = "developing"
	}
	app := bibliotheca.App{
		Debug:  env.Env == "dev",
		Commit: commit,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", env.Port),
		Handler:      app.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	eg.Go(func() error {
		return server.ListenAndServe()
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
