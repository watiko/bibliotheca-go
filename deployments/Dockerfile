FROM golang:1.13.7-buster as build

ENV AIR_VERSION 1.12.0
ENV TASK_VERSION 2.8.0

RUN apt update \
  && apt upgrade --yes \
  && apt clean \
  && rm --recursive --force /var/lib/apt/lists/*

RUN wget -O /usr/local/bin/air \
  https://github.com/cosmtrek/air/raw/v${AIR_VERSION}/bin/linux/air \
  && chmod +x /usr/local/bin/air

RUN wget -O task_linux_amd64.deb \
  https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_amd64.deb \
  && apt install ./task_linux_amd64.deb \
  && rm task_linux_amd64.deb

WORKDIR /app

COPY ./go.mod .
COPY ./go.mod .

RUN go mod download
