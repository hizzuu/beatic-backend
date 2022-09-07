#!/bin/bash
go install github.com/golang/mock/mockgen@latest
go generate ./...
go mod tidy
