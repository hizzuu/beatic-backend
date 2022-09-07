#!/bin/bash
go get -u github.com/99designs/gqlgen
go run github.com/99designs/gqlgen generate
go mod tidy
