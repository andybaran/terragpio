#! /bin/sh

cd cmd/client
GOOS=linux GOARCH=arm go build -v ./...
cd ../server
GOOS=linux GOARCH=arm go build -v ./...
cd ../..
scp ./cmd/server/server pi@10.15.21.124:/home/pi/bins
scp ./cmd/client/client pi@10.15.21.124:/home/pi/bins
