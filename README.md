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
$ task up
```

This example uses port number 8080. If you want use another port number, try `task up APP_PORT=18080`.

### Building

```console
$ task build:container
```
