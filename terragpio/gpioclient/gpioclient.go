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

func (c *Client) SetPWM(args SetPWMArgs) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	actualsettings, err := c.c.SetPWM(ctx, &pb.PWMRequest{Pin: args.Pin, Dutycycle: args.DutyCycle, Frequency: args.Freq})
	if err != nil {
		log.Fatalf("Error : ", c, err)
		return err
	}
	fmt.Println(actualsettings)
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
