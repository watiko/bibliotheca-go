#!/usr/bin/env bash

exec migrate -source file://scripts/migrations -database "${APP_DB_URL}" "$@"
