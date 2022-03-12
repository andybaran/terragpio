package main

import (
	"flag"
	"fmt"
	"net"

	pb "github.com/andybaran/fictional-goggles/terragpio"
	"google.golang.org/grpc"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/bmxx80"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

/*
	Common vars for use in authenticaiton if needed; currently not using.
	ToDo: Maybe make it possible to do this with vault?
*/
var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
)

/*
	Struct to represet a GPIO pin.
*/
type pinState struct {
	DutyCycle      *gpio.Duty
	Frequency      *physic.Frequency
	I2Caddr        *uint64
	I2Cbus         *string
	I2CDeviceOnBus *bmxx80.Dev
}

/*
Our server with a map to represent our pins
*/
type terragpioserver struct {
	pb.UnimplementedSetgpioServer
	Pins map[string]pinState //Our key is a string because it is the GPIO pin identified by name, ie: 'GPIO13'
}

func newServer() *terragpioserver {
	s := &terragpioserver{}
	s.Pins = make(map[string]pinState)
	return s
}

func main() {
	flag.Parse()
	host.Init()

	fmt.Printf("Pi? %v \n\n", rpi.Present())
	fmt.Printf("Available Pins: %+v \n\n", gpioreg.All())
	fmt.Printf("I2C Busses: %+v \n\n", i2creg.All())

	/*d, de := gpio.ParseDuty("90%")
	if de != nil {
		fmt.Println("duty cycle parsing error: ", de)
	}
	setPWMDutyCycle(d, physic.Frequency(25000), gpioreg.ByName(("GPIO13")))*/

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())

	// Listen for client connections
	grpcServer.Serve(lis)
}
