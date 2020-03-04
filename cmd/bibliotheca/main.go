package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca"
)

var commit string

type Env struct {
	Port  int    `default:"8080"`
	Env   string `default:"dev"`
	DbURL string `required:"true" split_words:"true"`
}

func main() {
	var env Env

	err := envconfig.Process("APP", &env)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	db, err := sqlx.Open("postgres", env.DbURL)
	if err != nil {
		log.Fatalf("failed to create db object: %v", err)
	}

	if commit == "" {
		commit = "developing"
	}

	app := bibliotheca.NewApp(env.Env, commit, db)

	var eg errgroup.Group

	eg.Go(func() error {
		return app.Run(env.Port)
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
