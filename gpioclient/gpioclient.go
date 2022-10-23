package gpioclient

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/andybaran/terragpio"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	TimeInterval    uint64
	BME280DevicePin string
	TemperatureMax  uint64
	TemperatureMin  uint64
	FanDevice       string
	DutyCycleMax    uint64
	DutyCylceMin    uint64
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

func (c *Client) SetBME280(args SetBME280Args) (*pb.PinSetResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.c.SetBME280(ctx, &pb.BME280Request{I2Cbus: args.I2CBus, I2Caddr: args.I2CAddr})
	if err != nil {
		log.Fatalf("Error : ", c, err)
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil
}

func (c *Client) StartFanController(args StartFanControllerArgs) (*pb.FanControllerResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.c.PWMDutyCycleOutput_BME280TempInput(ctx,
		&pb.FanControllerRequest{
			TimeInterval:    args.TimeInterval,
			BME280DevicePin: args.BME280DevicePin,
			TemperatureMax:  args.TemperatureMax,
			TemperatureMin:  args.TemperatureMin,
			FanDevicePin:    args.FanDevice,
			DutyCycleMax:    args.DutyCycleMax,
			DutyCycleMin:    args.DutyCylceMin})
	if err != nil {
		log.Fatalf("Error : ", c, err)
		return nil, err
	}
	return resp, nil
}

func NewClient(serverAddr string) (*Client, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)

	//conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewSetgpioClient(conn)

	return &Client{c: c}, nil

}
