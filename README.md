# bibliotheca-go

A WIP API backend for [bibliotheca-pwa](https://github.com/opt-tech/bibliotheca-pwa).

## Getting Started

### Prerequisites

- [task](https://github.com/go-task/task)
- Docker (supports BuildKit)
  - `docker`
  - `docker-compose`
- Go

### Developing

```console
$ task setup
$ task dc:up
```

This example uses port number 8080. If you want use another port number, try `env APP_PORT=18080 task dc:up`.

#### Database

Use psql in docker container.

```console
$ task dc:psql
```

If you want to connect to the development database from outside the container. Try this:

```console
$ psql "$(task dc:show-db-connection-string)"
```

### Building

```console
$ task build:container
```
