#!/usr/bin/env bash

FMT="*.go"

echo "Running tests..."
go test -v -cover -cpu 1,2,4 ./...
go test -v -cover -cpu 1,2,4 -race ./...

echo "Checking gofmt..."
fmt_res=$(gofmt -l -s $FMT)
if [ -n "${fmt_res}" ]; then
  echo -e "gofmt checking failed:\n${fmt_res}"
  exit 1
fi

echo "Checking govet..."
vet_res=$(go vet ./...)
if [ -n "${vet_res}" ]; then
  echo -e "govet checking failed:\n${vet_res}"
  exit 1
fi

echo "Success"
