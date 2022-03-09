#!/bin/bash

mkdir -p build/

go build -o ./build/producer ./cmd/producer.go
go build -o ./build/consumer ./cmd/consumer.go
go build -o ./build/check ./cmd/check.go
go build -o ./build/repair ./cmd/repair.go

cp -f ./build/producer ../eth-go-build/rinkeby/
cp -f ./build/producer ../eth-go-build/moonriver/

cp -f ./build/consumer ../eth-go-build/rinkeby/
cp -f ./build/consumer ../eth-go-build/moonriver/

cp -f ./build/repair ../eth-go-build/rinkeby/
cp -f ./build/repair ../eth-go-build/moonriver/
