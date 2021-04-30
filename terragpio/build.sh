#! /bin/sh

GOOS=linux GOARCH=arm go build ./...
scp ./cmd/server/server pi@10.15.21.124:/home/pi/bins
scp ./cmd/client/client pi@10.15.21.124:/home/pi/bins
