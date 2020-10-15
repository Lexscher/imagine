#!/bin/bash

set -e

echo " . . . . . . . . migrating imagine db 🧘‍♀️"
go run cmd/dbmigrate/main.go

echo " . . . . . . . . migrating imagine test db 🧪"
go run cmd/dbmigrate/main.go -dbname=imaginetest

echo " . . . . . . . . downloading CompileDaemon from github.com/githubnemo 🔻"
GO111MODULE=off go get github.com/githubnemo/CompileDaemon

echo " . . . . . . . . starting daemon 🐅 🦅"
CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main