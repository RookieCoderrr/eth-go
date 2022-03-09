#!/bin/bash

mkdir -p build/

go build -o ./build/producer ./cmd/producer.go
go build -o ./build/consumer ./cmd/consumer.go
go build -o ./build/check ./cmd/check.go
go build -o ./build/repair ./cmd/repair.go
