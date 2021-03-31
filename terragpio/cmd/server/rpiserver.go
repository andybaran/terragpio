package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "terragpio"

	"google.golang.org/grpc"
	"periph.io/x/conn/gpio/gpioreg/v3"
	"periph.io/x/conn/gpio/v3"
	"periph.io/x/conn/physic/v3"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
)

type terragpioserver struct {
	pb.UnimplementedSetgpioServer
}

// TODO : What should I return here? the whole Pwm seems unnecesary
// TODO : Review use of pointers and refs here..should I be using & in the body?
func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.Pwm) (*pb.Pwm, error) {
	//settings := settings
	pin := gpioreg.ByName(settings.Pin)
	if err := pin.PWM(gpio.Duty(settings.Dutycycle), physic.Frequency(settings.Frequency)); err != nil {
		println(err)
		return settings, err
	}
	return settings, nil
}

func newServer() *terragpioserver {
	s := &terragpioserver{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
