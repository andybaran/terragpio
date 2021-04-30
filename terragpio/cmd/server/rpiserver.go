package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/andybaran/fictional-goggles/terragpio"

	"google.golang.org/grpc"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
)

type pinState struct {
	DutyCycle gpio.Duty
	Frequency physic.Frequency
}
type terragpioserver struct {
	pb.UnimplementedSetgpioServer
	Pins map[string]pinState
}

func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.PWMRequest) (*pb.PWMResponse, error) {

	fmt.Printf("settings: %+v \n", settings)
	pin := gpioreg.ByName(settings.Pin)
	fmt.Printf("pin: %+v", pin)

	d, err := gpio.ParseDuty(settings.Dutycycle)
	if err != nil {
		println(err)
		return nil, err
	}

	var f physic.Frequency
	if err := f.Set(settings.Frequency); err != nil {
		println(err)
		return nil, err
	}

	if err := pin.PWM(d, f); err != nil {
		println(err)
		return nil, err
	}

	thisPinState := pinState{
		DutyCycle: d,
		Frequency: f,
	}

	s.Pins[settings.Pin] = thisPinState

	resp := pb.PWMResponse{Verified: true}
	return &resp, nil
}

func newServer() *terragpioserver {
	s := &terragpioserver{}

	//s.Pins = make(map[string]pinState)
	return s
}

func (s *terragpioserver) genPWMResponse() (response pb.PWMResponse) {

	var err string
	err = "notYet"
	//what's special about "nil"?

	if err != "notYet" {
		response.Verified = false
		return response
	}

	response.Verified = true
	return response

}

func main() {
	flag.Parse()
	host.Init()
	fmt.Printf("Available Pins: %+v \n", gpioreg.All())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}
