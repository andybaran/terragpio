clear
echo "*********** Reticulating Splines ***********"
echo "Building"
GOOS=linux GOARCH=arm go build -o ./server ./cmd/server/rpiserver.go 
GOOS=linux GOARCH=arm go build -o ./client ./cmd/client/rpiclient.go    
echo "Copying"
scp ./server pi@10.15.21.124:/home/pi/bins
scp ./client pi@10.0.0.157:/home/pi/bins
echo "***********     All Done        ***********"
