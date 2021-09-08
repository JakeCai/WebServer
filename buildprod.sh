#! /bin/bash

# Build web and other services

cd ~/work/src/github.com/jakecai/WebServer/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ~/work/src/github.com/jakecai/WebServer/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/work/src/github.com/jakecai/WebServer/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd ~/work/src/github.com/jakecai/WebServer/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web