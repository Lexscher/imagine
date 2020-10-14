#!/bin/bash

go run cmd/dbmigrate/main.go \
  -migrate=down \
  -dbname=imaginetest \
  -dbhost=localhost