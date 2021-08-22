package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/andybaran/fictional-goggles/terragpio"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10001", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func setPWM(client pb.SetgpioClient, settings *pb.PWMRequest) {
	fmt.Printf("Setting PWM --> %+v \n", settings)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	actualsettings, err := client.SetPWM(ctx, settings)
	if err != nil {
		log.Fatalf("Error : ", client, err)
	}
	fmt.Println(actualsettings)
}

func SetBME280(client pb.SetgpioClient, settings *pb.BME280Request) {
	fmt.Printf("Setting BME280 --> %+v \n", settings)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	actualsettings, err := client.SetBME280(ctx, settings)
	if err != nil {
		log.Fatalf("Error : ", client, err)
	}
	fmt.Println(actualsettings)
}

func main() {

	pinPtr := flag.String("pin", "", "GPIO Pin")
	dutyCyclePtr := flag.String("dutycycle", "", "Duty cycle")
	freqPtr := flag.String("frequency", "", "Frequency")

	I2Cbus := flag.String("I2Cbus", "1", "I2C Bus")        //Very likely "1" on a raspberry pi
	I2Caddr := flag.Uint64("I2Caddr", 0x77, "I2C Address") //BME280 may also be 0x76

	flag.Parse()

	var opts []grpc.DialOption

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
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSetgpioClient(conn)

	// Set PWM
	setPWM(client, &pb.PWMRequest{
		Pin:       *pinPtr,       //"GPIO13",
		Dutycycle: *dutyCyclePtr, //"100%",
		Frequency: *freqPtr,      //"25000",
	})

	SetBME280(client, &pb.BME280Request{
		I2Cbus:  *I2Cbus,
		I2Caddr: *I2Caddr, // "0x76"
	})

}
