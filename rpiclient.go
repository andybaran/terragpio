package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "terragpio"

	"google.golang.org/grpc"
)

var (
	//tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	//caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	//serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// printFeature gets the feature for the given point.
func setPWM(client pb.SetgpioClient, settings *pb.Pwm) {
	log.Printf("Setting PWM (%d, %d, %d)", settings.Pin, settings.Dutycycle, settings.Frequency)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.SetPWM(ctx, settings)
	if err != nil {
		log.Fatalf("Error : ", client, err)
	}
	log.Println(settings)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	/*
		if *tls {
			if *caFile == "" {
				*caFile = data.Path("x509/ca_cert.pem")
			}
			creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
			if err != nil {
				log.Fatalf("Failed to create TLS credentials %v", err)
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}
		opts = append(opts, grpc.WithInsecure())

		opts = append(opts, grpc.WithBlock())
	*/
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSetgpioClient(conn)

	// Set PWM
	setPWM(client, &pb.Pwm{
		Pin:       "GPIO6",
		Dutycycle: "100%",
		Frequency: "25M",
	})

}
