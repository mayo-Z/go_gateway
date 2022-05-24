#!/bin/sh

export GO111MODULE=on  && export GOPROXY=https://goproxy.cn && go mod tidy
GOOS=linux GOARCH=amd64 go build -o ./bin/go-gateway
docker build -f dockerfile-dashboard -t go-gateway-dashboard .
docker build -f dockerfile-server -t go-gateway-server .

