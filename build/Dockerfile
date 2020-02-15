# syntax=docker/dockerfile:1.1.3-experimental
# https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/experimental.md

ARG AIR_VERSION=1.12.0
ARG TASK_VERSION=2.8.0
ARG MIGRATE_VERSION=4.8.0


FROM golang:1.13.7-buster as base

RUN apt update \
  && apt upgrade --yes \
  && apt clean \
  && rm --recursive --force /var/lib/apt/lists/*

ARG TASK_VERSION
RUN wget -O task_linux_amd64.deb \
  https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_amd64.deb \
  && apt install ./task_linux_amd64.deb \
  && rm task_linux_amd64.deb

ARG MIGRATE_VERSION
RUN wget https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
  && tar xf migrate.linux-amd64.tar.gz \
  && mv migrate.linux-amd64 /usr/local/bin/migrate

WORKDIR /app

COPY ./go.mod .
COPY ./go.mod .

RUN go mod download


FROM base as dev

RUN apt update \
  && apt install -y postgresql-client-11 \
  && apt clean \
  && rm --recursive --force /var/lib/apt/lists/*

ARG AIR_VERSION
RUN wget -O /usr/local/bin/air \
  https://github.com/cosmtrek/air/raw/v${AIR_VERSION}/bin/linux/air \
  && chmod +x /usr/local/bin/air


FROM base as build

# https://github.com/golang/go/issues/26492#issuecomment-435462350
ENV GOFLAGS "-tags=netgo -ldflags=-extldflags=-static"
ENV GOCACHE /.cache

RUN \
  --mount=type=bind,rw,target=/app,source=. \
  --mount=type=cache,target=/.cache \
  task build \
  && mkdir -p /tmp/app/bin \
  && cp -R /app/bin /tmp/app

FROM gcr.io/distroless/base-debian10 as app

COPY --from=build /tmp/app/bin/ /app/bin/

CMD ["/app/bin/bibliotheca"]

EXPOSE 8080


# use busybox as shell
FROM gcr.io/distroless/base-debian10:debug as migration

COPY ./scripts/migrations /app/scripts/migrations
COPY ./scripts/migrate.sh /app/scripts/
COPY --from=base /usr/local/bin/migrate /app/migrate

ENV PATH "/app:$PATH"
WORKDIR /app

CMD ["/app/scripts/migrate.sh", "up"]
