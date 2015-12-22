#!/usr/bin/env bash
GOOS=linux GOARCH=arm go build \
  -a \
  -tags netgo \
  -installsuffix netgo \
  -ldflags '-w -s' \
  -o rasp3 \
  main.go
