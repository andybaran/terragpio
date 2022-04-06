package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
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
	DutyCycle      *gpio.Duty
	Frequency      *physic.Frequency
	I2Caddr        *uint64
	I2Cbus         *string
	I2CDeviceOnBus *bmxx80.Dev
}

/*
Our server with a map to represent our pins
*/
type terragpioserver struct {
	pb.UnimplementedSetgpioServer
	Pins map[string]pinState //Our key is a string because it is the GPIO pin identified by name, ie: 'GPIO13'
}

func newServer() *terragpioserver {
	s := &terragpioserver{}
	s.Pins = make(map[string]pinState)
	return s
}

// Set frequency and duty cycle on a pin
func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.PWMRequest) (*pb.PinSetResponse, error) {

	//fmt.Printf("settings: %+v \n\n", settings)

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
		println(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Failed setting duty cycle and/or frequency, %s", err))
	}

	thisPinState := pinState{
		DutyCycle: &d,
		Frequency: &f,
	}

	s.Pins[settings.Pin] = thisPinState

	fmt.Printf("Duty Cycle: %+v \n", d)
	fmt.Printf("Frequency : %+v \n", f)

	resp := pb.PinSetResponse{PinNumber: settings.Pin}
	return &resp, nil
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

func (s *terragpioserver) SetBME280(ctx context.Context, settings *pb.BME280Request) (*pb.PinSetResponse, error) {
	bus, err := i2creg.Open(settings.I2Cbus)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to open the i2c bus, %s", err))
	}
	//defer bus.Close()

	dev, err := bmxx80.NewI2C(bus, uint16(settings.I2Caddr), &bmxx80.DefaultOpts)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to intitialize your bmxx80 i2c device, %s", err))
	}
	//defer dev.Halt()

	thisPinState := pinState{
		I2Caddr:        &settings.I2Caddr,
		I2Cbus:         &settings.I2Cbus,
		I2CDeviceOnBus: dev,
	}

	//for an i2c device we track it using it's bus and address on that bus instead of the specific pin
	i2cBusAddr := settings.I2Cbus
	i2cBusAddr += strconv.FormatUint(settings.I2Caddr, 10)
	s.Pins[i2cBusAddr] = thisPinState
	resp := pb.PinSetResponse{PinNumber: i2cBusAddr}
	return &resp, nil
}

// Return temperature, pressure and humidity readings from a BME280 sensor connected via i2c
func (s *terragpioserver) SenseBME280(ctx context.Context, pin *pb.PinSetRequest) (*pb.BME280Response, error) {

	// Read temperature from the sensor:
	var env physic.Env
	println("in the sensing function")
	println(s.Pins[pin.PinNumber].I2CDeviceOnBus)
	dev := s.Pins[pin.PinNumber].I2CDeviceOnBus
	if err := dev.Sense(&env); err != nil {
		println("error sensing")
		log.Fatal(err)
	}
	fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)

	resp := pb.BME280Response{Temperature: env.Temperature.String(), Pressure: env.Pressure.String(), Humidity: env.Humidity.String()}
	return &resp, nil
}

// Set duty cycle on a pin based on the temperature reading from a BME280
func (s *terragpioserver) PWMDutyCycleOutput_BME280TempInput(ctx context.Context, settings *pb.FanControllerRequest) (*pb.FanControllerResponse, error) {
	//setup the PWM device
	//s.SetPWM(ctx, settings.FanDevice)
	println("entered calculation")
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
			println("ticker started; trying to sense from bme280")
			fmt.Println("BME280DevicePin: ", settings.BME280DevicePin)
			r, err := s.SenseBME280(ctx, &pb.PinSetRequest{PinNumber: settings.BME280DevicePin})
			if err != nil {
				println("error sensing bme280")
				panic(err)
			}

			/* Temperature value will be returned as a string like "10 C"
			*  convert it to a physic.Temperature so we can convert it to a uint64 and do some math
			 */
			println("setting duty cycle")
			t.Set(r.Temperature)
			fmt.Println("temp: ", t)
			fmt.Println("t in celsius", uint64(t.Celsius()))
			fmt.Println("temp max: ", strconv.FormatUint(settings.TemperatureMax, 10))
			fmt.Println("duty cycle max: ", strconv.FormatUint(settings.DutyCycleMax, 10))
			fmt.Println("slope: ", strconv.FormatUint(slope, 10))
			fmt.Println("")

			floatd := 0 - (float64(slope)*(float64(settings.TemperatureMax)-(t.Celsius())) - float64(settings.DutyCycleMax))

			fmt.Println("floatd: ", floatd)
			intd := int(math.Round(floatd))
			fmt.Println("inttd: ", intd)
			stringd := strconv.Itoa(intd) + "%"
			fmt.Println("stringd: ", stringd)

			d, err := gpio.ParseDuty(stringd)
			if err != nil {
				fmt.Println("error parsing duty cycle? : ", err)
				panic(err)
			}
			//d := settings.DutyCycleMax - (slope * (uint64(t.Celsius())))

			//set the dutycycle
			fan := s.Pins[settings.FanDevicePin]
			f.Set(fan.Frequency.String())
			setPWMDutyCycle(d,
				f,
				gpioreg.ByName(settings.FanDevicePin))
		}
	}()

	resp := pb.FanControllerResponse{}
	return &resp, nil
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
