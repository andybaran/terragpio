clear
echo "*********** Reticulating Splines ***********"
echo "Building"
GOOS=linux GOARCH=arm go build -o ./server ./cmd/server/terragpioserver.go  
GOOS=linux GOARCH=arm go build -o ./client ./cmd/client/terragpioclient.go     
echo "Copying"
<<<<<<< HEAD
scp ./server pi@10.15.21.124:/home/pi/bins
scp ./client pi@10.15.21.124:/home/pi/bins
echo "Cleaning up"
rm ./server
rm ./client
=======
scp ./server pi@10.15.21.201:/home/pi/bins
scp ./client pi@10.15.21.201:/home/pi/bins
>>>>>>> refs/remotes/origin/main
echo "***********     All Done        ***********"
