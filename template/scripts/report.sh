#!/bin/bash

ERR="false"

run () {
    echo "Running ${@:1}"
    diff -u <(echo -n) <(${@:1})
    if [ $? -ne 0 ]
    then
        ERR="true"
    fi
}

run gofmt -d .
run go vet ./...
run golint ./...
run gocyclo -over 15 .
run ineffassign .
run misspell .

if [ $ERR == "true" ]
then
    exit 1
fi