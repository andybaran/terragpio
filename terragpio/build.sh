#! /bin/sh

# cd cmd/client
# GOOS=linux GOARCH=arm go build -v ./...
# cd ../server
echo "************************************************"
echo "Building"
echo "************************************************"
echo ""
GOOS=linux GOARCH=arm go build -v -a ./...
GOOS=linux GOARCH=arm go build -v -a ./cmd/server/rpiserver.go -o ./server
GOOS=linux GOARCH=arm go build -v -a ./cmd/client/rpiclient.go -o ./client    
echo ""
echo "************************************************"
echo "Done building"
echo "************************************************"
echo ""
echo "************************************************"
echo "Copying"
echo "************************************************"
echo ""
cd ../..
scp ./cmd/server/server pi@10.15.21.124:/home/pi/bins
scp ./cmd/client/client pi@10.15.21.124:/home/pi/bins
echo ""
echo "************************************************"
echo "All done"
echo "************************************************"
echo ""