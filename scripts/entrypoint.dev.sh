#!/bin/bash

set -e

go run cmd/dbmigrate/main.go

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main
