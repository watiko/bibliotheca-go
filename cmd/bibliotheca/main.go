package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca"
	"golang.org/x/sync/errgroup"
)

type Env struct {
	Port int `default:"8080"`
}

func main() {
	var env Env

	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	var eg errgroup.Group

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", env.Port),
		Handler:      bibliotheca.Router(),
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
