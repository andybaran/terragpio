package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "terragpio"

	"google.golang.org/grpc"
	"periph.io/x/conn/gpio"
	"periph.io/x/conn/gpio/gpioreg"
	"periph.io/x/conn/physic"
)

type terragpioserver struct {
	pb.UnimplementedSetgpioServer
}

// TODO : What should I return here?
func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.Pwm) (*pb.Pwm, error) {
	//settings := settings
	pin := gpioreg.ByName(settings.Pin)
	realDutyCycle := gpio.DutyMax * settings.Dutycycle
	realFrequency := settings.Frequency * physic.Hertz
	if err := pin.PWM(realDutyCycle, realFrequency); err != nil {
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
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
