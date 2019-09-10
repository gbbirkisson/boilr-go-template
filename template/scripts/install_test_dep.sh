#!/bin/bash


# This is done so go modules do not add those
# projects as dependencies to this project
TMP="$(mktemp -d)"
cd "${TMP}"

go get -v -u \
    github.com/client9/misspell/cmd/misspell \
    github.com/fzipp/gocyclo \
    github.com/gordonklaus/ineffassign \
    golang.org/x/lint/golint