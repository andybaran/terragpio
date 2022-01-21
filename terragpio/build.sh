clear
echo "*********** Reticulating Splines ***********"
echo "Building"
<<<<<<< HEAD
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
=======
GOOS=linux GOARCH=arm go build -o ./server ./cmd/server/rpiserver.go 
GOOS=linux GOARCH=arm go build -o ./client ./cmd/client/rpiclient.go    
>>>>>>> bme280
echo "Copying"
scp ./server pi@10.15.21.124:/home/pi/bins
scp ./client pi@10.15.21.124:/home/pi/bins
echo "***********     All Done        ***********"
