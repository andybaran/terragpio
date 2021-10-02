package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

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

//ToDo: Add a i2C state struct

type terragpioserver struct {
	pb.UnimplementedSetgpioServer
	Pins map[string]pinState
}

func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.PWMRequest) (*pb.PWMResponse, error) {

	fmt.Printf("settings: %+v \n\n", settings)
	pin := gpioreg.ByName(settings.Pin)
	fmt.Printf("pin: %+v \n\n", pin)

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

	fmt.Printf("Duty Cycle: %+v \n", d)
	fmt.Printf("Frequency : %+v \n", f)

	resp := pb.PWMResponse{Verified: true}
	return &resp, nil
}

func (s *terragpioserver) GetBME280(ctx context.Context, settings *pb.BME280Request) (*pb.BME280Response, error) {
	fmt.Printf("settings: %+v \n\n", settings)

	bus, err := i2creg.Open(settings.I2Cbus)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer bus.Close()

	dev, err := bmxx80.NewI2C(bus, uint16(settings.I2Caddr), &bmxx80.DefaultOpts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer dev.Halt()

	// Read temperature from the sensor:
	var env physic.Env
	if err = dev.Sense(&env); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)

	resp := pb.BME280Response{Temperature: env.Temperature.String(), Pressure: env.Pressure.String(), Humidity: env.Humidity.String()}
	return &resp, nil
}

func newServer() *terragpioserver {
	s := &terragpioserver{}
	s.Pins = make(map[string]pinState)
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

func readBME() (physic.Temperature, error) {
	bus, err := i2creg.Open("") //just open the first bus found
	fmt.Println("i2c bus opened")
	if err != nil {
		return 0, err
	}
	defer bus.Close()

	dev, err := bmxx80.NewI2C(bus, uint16(0x77), &bmxx80.DefaultOpts) //0x77 is default for the bme280 I currently have
	fmt.Println("ready to read values")
	if err != nil {
		return 0, err
	}
	defer dev.Halt()

	// Read temperature from the sensor:
	var env physic.Env
	if err = dev.Sense(&env); err != nil {
		return 0, err
	}
	fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)
	fmt.Println("returning env")
	return env.Temperature, nil
}

func setPWMDutyCycle(d gpio.Duty, f physic.Frequency, p gpio.PinIO) error {

	if err := p.PWM(d, f); err != nil {
		println(err)
		return err
	}
	println("duty cycle % = ", d)
	return nil

}

func main() {
	flag.Parse()
	host.Init()

	fmt.Printf("Pi? %v \n\n", rpi.Present())
	fmt.Printf("Available Pins: %+v \n\n", gpioreg.All())
	fmt.Printf("I2C Busses: %+v \n\n", i2creg.All())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())

	// Listen for client connections
	go func() {
		grpcServer.Serve(lis)
	}()

	// Make some channels
	temperatureChan := make(chan physic.Temperature)
	calculateOutput := make(chan physic.Temperature)
	//	setDutyCycle := make(chan gpio.Duty)

	go func() {
		for t := range temperatureChan {
			calculateOutput <- t
			fmt.Println("Recieved temperature of: ", t)
		}
	}()

	// Calculate curve
	//// I need a struct here that can pass in the max's and min's
	go func() {
		for c := range calculateOutput {
			//// temperature range in celsius (x)
			var tMax int = 35
			var tMin int = 5

			//// might as well make duty cycle configurable too (y)
			var dMax int = 100
			var dMin int = 20

			//calculate our slope
			s := (dMax - tMax) / (dMin - tMin)

			d := (s*(tMax) - int(c.Celsius()) + dMax)
			//calculate duty cycle (y axis using y = mx+b)
			setPWMDutyCycle(gpio.Duty(d),
				25000,
				gpioreg.ByName("GPIO13"))
			//setDutyCycle <- gpio.Duty((s*(tMax) - int(c.Celsius()) + dMax))
		}
	}()

	/*go func() {
		d := string(setDutyCycle)
		if err != nil {
			println(err)
		}
		setPWMDutyCycle(d, 25000, gpioreg.ByName("GPIO13"))
	}()*/

	/* Loop every 2 seconds
	Read temp from readBME()
	Write temp to temperatureChan
	*/
	myTicker := time.NewTicker(time.Second * 2)
	for range myTicker.C {
		actualTemp, err := readBME()
		if err != nil {
			panic(err)
		}
		temperatureChan <- actualTemp
		println("sending actual temp")
	}

}

/* 1) readBME get temperature read working (i2c bus, etc)
2) send value to temperatureChan
3) go routine parallel to read from temperatureChan...start this before time loop...needs to be go routine not just a func
4) read and log value
5) then how do we transform that to a fan speed
6) set fan speed
*/
