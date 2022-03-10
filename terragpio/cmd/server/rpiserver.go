package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/andybaran/fictional-goggles/terragpio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/bmxx80"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

/*
	Common vars for use in authenticaiton if needed; currently not using.
	ToDo: Maybe make it possible to do this with vault?
*/
var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
)

/*
	Struct to represet a GPIO pin.
*/
type pinState struct {
	DutyCycle *gpio.Duty
	Frequency *physic.Frequency
	I2Caddr *uint64
	I2Cbus *string
	I2CDeviceOnBus *bmxx80	
}

/*
Our server with A map to represent our pins
*/
type terragpioserver struct {
	pb.UnimplementedSetgpioServer
	Pins map[string]pinState //Our key is a string because it is the GPIO pin identified by name, ie: 'GPIO13'
}

// Set frequency and duty cycle on a pin
func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.PWMRequest) (*pb.genericPinSetResponse, error) {

	//fmt.Printf("settings: %+v \n\n", settings)
	// ToDo: How in the heck do I handle an error here? Just try to catch invalid input?
	pin := gpioreg.ByName(settings.Pin) 

	d, err := gpio.ParseDuty(settings.Dutycycle)
	if err != nil {
		println(err)
		return nil, err
	}

	var f physic.Frequency
	if err := f.Set(settings.Frequency); err != nil {
		// println(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to compute the frequency, %s", err))
		// Unknown seems like the most logical since this isn't really a typical API response that we're expecting from the GPIO library
	}

	if err := pin.PWM(d, f); err != nil {
		// println(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Failed setting duty cycle and/or frequency, %s", err))
	}

	thisPinState := pinState{
		DutyCycle: d,
		Frequency: f,
	}

	s.Pins[settings.Pin] = thisPinState

	fmt.Printf("Duty Cycle: %+v \n", d)
	fmt.Printf("Frequency : %+v \n", f)

	resp := pb.PWMResponse{pinNumber: settings.Pin}
	return &resp, nil
}

func (s *terragpioserver) SetBME280(ctx context.Context, settings *pb.BME280Request) (*pb.genericPinSetResponse, error) {
	bus, err := i2creg.Open(settings.I2Cbus)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to open the i2c bus, %s", err))
	}
	defer bus.Close()

	dev, err := bmxx80.NewI2C(bus, uint16(settings.I2Caddr), &bmxx80.DefaultOpts)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to intitialize your bmxx80 i2c device, %s", err))
	}
	defer dev.Halt()
	
	thisPinState := pinState{
		I2Caddr: settings.I2Caddr,
		I2Cbus: settings.I2Cbus,
		I2CDeviceOnBus: dev,
	}

	s.Pins[settings.Pin] = thisPinState
	resp := pb.genericPinSetResponse{pinNumber: string(thisPinState.I2Caddr)}
	return &resp, nil
}

// Return temperature, pressure and humidity readings from a BME280 sensor connected via i2c
func (s *terragpioserver) SenseBME280(ctx context.Context, pin *pb.genericPin) (*pb.BME280Response, error) {
	
	// Read temperature from the sensor:
	var env physic.Env
	dev = s.Pins[pin.pin].I2CDeviceOnBus
	if err = dev.Sense(&env); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)

	resp := pb.BME280Response{Temperature: env.Temperature.String(), Pressure: env.Pressure.String(), Humidity: env.Humidity.String()}
	return &resp, nil
}

// Set duty cycle on a pin based on the temperature reading from a BME280
func (s *terragpioserver) PWMDutyCycleOutput_BME280TempInput(ctx context.Context, settings *pb.FanControllerRequest) (*pb.FanController, error) {
	//setup the PWM device
	//s.SetPWM(ctx, settings.FanDevice)

	/* Calculate slope so we that when given max and min duty cycle settings and temperature readings.
	*  We use this to calculate duty cycle (d) based on temperature readings (r.Temperature).
	 */
	slope := (settings.DutyCycleMax - settings.DutyCycleMin) / (settings.TemperatureMax - settings.TemperatureMin)

	//Setup the temperature and frequency vars so we can use them later
	var t physic.Temperature
	var f physic.Frequency

	/* We want to start a loop here that gets the temp and sets the duty cycle
	*  However, we don't want to be in a blocking loop so the loop can be brought into a go routine
	 */

	//Convert the uint64 value in setttings.TimeInterval to time.Duration so we can convert to time.Second and use it as the input for our ticker
	dutyCycleTicker := time.NewTicker(time.Second * time.Duration(settings.TimeInterval))

	//start the ticker
	go func() {
		for range dutyCycleTicker.C {
			//read the values from the BME280
			r, err := s.SenseBME280(ctx, pb.genericPin{Pin = settings.BME280DevicePin})
			if err != nil {
				panic(err)
			}

			/* Temperature value will be returned as a string like "10 C"
			*  convert it to a physic.Temperature so we can convert it to a uint64 and do some math
			 */
			t.Set(r.Temperature)
			d, err := gpio.ParseDuty(strconv.FormatUint(settings.DutyCycleMax-(slope*(uint64(t.Celsius()))), 10) + "%")
			if err != nil {
				//fmt.Println("error parsing duty cycle: ", err)
				panic(err)
			}
			//d := settings.DutyCycleMax - (slope * (uint64(t.Celsius())))

			//set the dutycycle
			fan = s.Pins[settings.fanDevicePin]
			f.Set(fan.Frequency)
			setPWMDutyCycle(d,
				f,
				gpioreg.ByName(settings.fanDevicePin))

		}
	}()

	resp := pb.FanControllerResponse{}
	return &resp, nil
}

func newServer() *terragpioserver {
	s := &terragpioserver{}
	s.Pins = make(map[string]pinState)
	return s
}


func setPWMDutyCycle(d gpio.Duty, f physic.Frequency, p gpio.PinIO) error {

	if err := p.PWM(d, f); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("pin for pwm: ", p)
	fmt.Println("duty cycle: ", d)
	fmt.Println("frequency: ", f)
	println()
	return nil
}

func main() {
	flag.Parse()
	host.Init()

	fmt.Printf("Pi? %v \n\n", rpi.Present())
	fmt.Printf("Available Pins: %+v \n\n", gpioreg.All())
	fmt.Printf("I2C Busses: %+v \n\n", i2creg.All())

	/*d, de := gpio.ParseDuty("90%")
	if de != nil {
		fmt.Println("duty cycle parsing error: ", de)
	}
	setPWMDutyCycle(d, physic.Frequency(25000), gpioreg.ByName(("GPIO13")))*/

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())

	// Listen for client connections
	grpcServer.Serve(lis)
}
