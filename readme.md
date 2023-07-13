
- line 247 in terragpioserver.go and line 17 in terragpioclient.go currently hardcode the ip address and port that the server will listen on and client will connect to; change this as needed before running build.sh

- build.sh will compile both the command line client and server binaries and push them via scp to a raspberry pi; you will need to change the ip address hard coded into the script for scp to succeed

- once pushed to the pi run the server binary with sudo otherwise you won't have proper access to the GPIO pins; the client binary does not need to be run with sudo