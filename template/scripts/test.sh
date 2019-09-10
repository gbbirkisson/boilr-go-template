#!/bin/sh

set -e

echo 'Running go_test...'
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...