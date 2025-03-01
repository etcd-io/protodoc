#!/usr/bin/env bash

echo "Running tests..."
go test -v -cover -cpu 1,2,4 ./...
go test -v -cover -cpu 1,2,4 -race ./...

echo "Checking gofmt..."
# We utilize 'go fmt' to find all files suitable for formatting,
# but reuse full power gofmt to perform just RO check.
fmt_res=$(go fmt -n ./... | sed 's| -w | -d |g' | sh)
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
