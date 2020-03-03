package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
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

	var eg errgroup.Group

	if commit == "" {
		commit = "developing"
	}
	app := bibliotheca.NewApp(env.Env, commit, env.DbURL)

	eg.Go(func() error {
		return app.Run(env.Port)
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
