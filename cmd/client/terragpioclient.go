package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/andybaran/terragpio"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("server_addr", "10.15.21.201:1234", "The server address in the format of host:port")
)

func setPWM(client pb.SetgpioClient, settings *pb.PWMRequest) (*pb.PinSetResponse, error) {
	fmt.Printf("Setting PWM --> %+v \n", settings)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SetPWM(ctx, settings)
	if err != nil {
		log.Fatalf("Error from setPWM: %s", err)
	}
	return resp, nil
}

func setBME280(client pb.SetgpioClient, settings *pb.BME280Request) {
	fmt.Printf("Setting BME280 --> %+v \n", settings)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	actualsettings, err := client.SetBME280(ctx, settings)
	if err != nil {
		log.Fatalf("Error from SetBME280: %w", err)
	}
	fmt.Println(actualsettings)
}

func startFanController(client pb.SetgpioClient, settings *pb.FanControllerRequest) {
	fmt.Printf("Setting up a fan controller --> %+v \n", settings)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	actualsettings, err := client.PWMDutyCycleOutput_BME280TempInput(ctx, settings)
	if err != nil {
		fmt.Printf("Error from StartFanController: %s", err)
	}
	fmt.Println(actualsettings)
}

func sensePWM(client pb.SetgpioClient, settings *pb.PinSetRequest) (string, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	actualsettings, err := client.SensePWM(ctx, settings)
	if err != nil {
		fmt.Printf("Error from SensePWM: %s", err)
	}
	fmt.Println(actualsettings)
	return actualsettings.Dutycycle, actualsettings.Frequency
}

func main() {

	// Flags to setup a PWM device on a GPIO pin, likely a fan
	pinPtr := flag.String("pin", "GPIO13", "GPIO Pin")
	dutyCyclePtr := flag.String("dutycycle", "10%", "Duty cycle")
	freqPtr := flag.String("frequency", "25000", "Frequency")

	// Flags to setup I2C bus and device. ie: a BME280 on bus 1
	I2Cbus := flag.String("I2Cbus", "1", "I2C Bus")          //Very likely "1" on a raspberry pi
	I2Caddr := flag.String("I2Caddr", "0x77", "I2C Address") //BME280 may also be 0x76

	// Flags to tie BME280 sensor and fan together
	timeInterval := flag.Uint64("timeInterval", 5, "Time in seconds")
	temperatureMax := flag.Uint64("temperatureMax", 100, "Max temp")
	temperatureMin := flag.Uint64("temperatureMin", 15, "Min temp")
	dutyCycleMax := flag.Uint64("dutyCycleMax", 100, "Max duty cycle")
	dutyCycleMin := flag.Uint64("dutyCycleMin", 10, "Min duty cycle")

	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial: ", *serverAddr, err)
	}
	defer conn.Close()
	client := pb.NewSetgpioClient(conn)

	// Set PWM
	setPWM(client, &pb.PWMRequest{
		Pin:       *pinPtr,       //"GPIO13",
		Dutycycle: *dutyCyclePtr, //"100%",
		Frequency: *freqPtr,      //"25000",
	})

	// SetupBME280 device
	setBME280(client, &pb.BME280Request{
		I2Cbus:  *I2Cbus,
		I2Caddr: *I2Caddr, // "0x76"
	})

	i2cBusAddr := *I2Cbus
	i2cBusAddr += *I2Caddr
	startFanController(client, &pb.FanControllerRequest{
		TimeInterval:    *timeInterval,
		BME280DevicePin: i2cBusAddr,
		TemperatureMax:  *temperatureMax,
		TemperatureMin:  *temperatureMin,
		FanDevicePin:    *pinPtr,
		DutyCycleMax:    *dutyCycleMax,
		DutyCycleMin:    *dutyCycleMin,
	})

}
