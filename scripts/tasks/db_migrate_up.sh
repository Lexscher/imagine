#!/bin/bash

go run cmd/dbmigrate/main.go \
  -migrate=up \
  -dbname=imaginetest \
  -dbhost=localhost