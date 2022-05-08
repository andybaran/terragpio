package gpioclient

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/andybaran/fictional-goggles/terragpio"

	"google.golang.org/grpc"
)

type Client struct {
	c pb.SetgpioClient
}

type SetPWMArgs struct {
	Pin       string
	DutyCycle string
	Freq      string
}

type SetBME280Args struct {
	I2CBus  string
	I2CAddr uint64
}

type StartFanControllerArgs struct {
	timeInterval   uint64
	BME280Device   pb.BME280Request
	temperatureMax uint64
	temperatureMin uint64
	fanDevice      pb.PWMRequest
	dutyCycleMax   string
	dutyCylceMin   string
}

func (c *Client) SetPWM(args SetPWMArgs) (*pb.PinSetResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.c.SetPWM(ctx, &pb.PWMRequest{Pin: args.Pin, Dutycycle: args.DutyCycle, Frequency: args.Freq})
	if err != nil {
		log.Fatalf("Error : ", c, err)
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil
}

func (c *Client) SetBME280(args SetBME280Args) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	actualsettings, err := c.c.SetBME280(ctx, &pb.BME280Request{I2Cbus: args.I2CBus, I2Caddr: args.I2CAddr})
	if err != nil {
		log.Fatalf("Error : ", c, err)
		return err
	}
	fmt.Println(actualsettings)
	return nil
}

func (c *Client) StartFanController(args StartFanControllerArgs) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.c.PWMDutyCycleOutput_BME280TempInput(ctx, &pb.FanControllerRequest{TimeInterval: args.timeInterval,
		BME280Device:   &args.BME280Device,
		TemperatureMax: args.temperatureMax,
		TemperatureMin: args.temperatureMin,
		FanDevice:      &args.fanDevice,
		DutyCycleMax:   args.dutyCycleMax,
		DutyCycleMin:   args.dutyCylceMin,
	})
	if err != nil {
		log.Fatalf("Error : ", c, err)
		return err
	}
	return nil
}

func NewClient(serverAddr string) (*Client, error) {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewSetgpioClient(conn)

	return &Client{c: c}, nil

}
