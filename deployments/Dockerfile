# syntax=docker/dockerfile:1.1.3-experimental
# https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/experimental.md

FROM golang:1.13.7-buster as base

RUN apt update \
  && apt upgrade --yes \
  && apt clean \
  && rm --recursive --force /var/lib/apt/lists/*

ENV TASK_VERSION 2.8.0
RUN wget -O task_linux_amd64.deb \
  https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_amd64.deb \
  && apt install ./task_linux_amd64.deb \
  && rm task_linux_amd64.deb

WORKDIR /app

COPY ./go.mod .
COPY ./go.mod .

RUN go mod download


FROM base as dev

ENV AIR_VERSION 1.12.0
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
