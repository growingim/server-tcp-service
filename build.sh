#!/usr/bin/env bash

go build -o ./build/server  main.go
go build -o ./build/client  ./client/client.go
