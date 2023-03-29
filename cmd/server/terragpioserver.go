package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	pb "github.com/andybaran/terragpio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/bmxx80"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

/*
Struct to represet a GPIO pin.
*/
type pinState struct {
	DutyCycle      *gpio.Duty
	Frequency      *physic.Frequency
	I2Caddr        *string
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

/* TODO: This really only be allowed to be set to proper PWM pins and give a warning if pin(s) other than those available for PWM
are passed. The gpio library may even just handle this for us? Regardless, for now we will assume a modern RPI is being used but develop a
framework that could incorporate other boards if needed.
*/

// Set frequency and duty cycle on a pin
func (s *terragpioserver) SetPWM(ctx context.Context, settings *pb.PWMRequest) (*pb.PinSetResponse, error) {

	pin := gpioreg.ByName(settings.Pin)

	d, err := gpio.ParseDuty(settings.Dutycycle)
	if err != nil {
		println(err)
		return nil, err
	}

	var f physic.Frequency
	if err := f.Set(settings.Frequency); err != nil {
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

	fmt.Printf("Initial Fan Duty Cycle: %+v \n", d)

	resp := pb.PinSetResponse{PinNumber: settings.Pin}
	return &resp, nil
}

func (s *terragpioserver) setPWMDutyCycle(d gpio.Duty, f physic.Frequency, p gpio.PinIO) error {

	if err := p.PWM(d, f); err != nil {
		fmt.Println(err)
		return err
	}

	thisPinState := pinState{
		DutyCycle: &d,
		Frequency: &f,
	}

	s.Pins[p.String()] = thisPinState

	fmt.Printf("Current Fan Duty Cycle: %+v \n", d)
	return nil
}

func (s *terragpioserver) SetBME280(ctx context.Context, settings *pb.BME280Request) (*pb.PinSetResponse, error) {
	bus, err := i2creg.Open(settings.I2Cbus)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to open the i2c bus, %s", err))
	}

	//Convert address from string to uint16
	i, err := strconv.ParseUint(settings.I2Caddr, 10, 16)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to convert i2c address from string to uint, %s", err))
	}

	dev, err := bmxx80.NewI2C(bus, uint16(i), &bmxx80.DefaultOpts)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("Unable to intitialize your bmxx80 i2c device, %s", err))
	}

	thisPinState := pinState{
		I2Caddr:        &settings.I2Caddr,
		I2Cbus:         &settings.I2Cbus,
		I2CDeviceOnBus: dev,
	}

	// We track an i2c device using it's bus and address on that bus instead of the specific pin
	i2cBusAddr := settings.I2Cbus
	i2cBusAddr += settings.I2Caddr
	s.Pins[i2cBusAddr] = thisPinState
	resp := pb.PinSetResponse{PinNumber: i2cBusAddr}
	return &resp, nil
}

// Return temperature, pressure and humidity readings from a BME280 sensor connected via i2c
func (s *terragpioserver) SenseBME280(ctx context.Context, pin *pb.PinSetRequest) (*pb.BME280Response, error) {

	// Read temperature from the sensor:
	var env physic.Env
	dev := s.Pins[pin.PinNumber].I2CDeviceOnBus
	if err := dev.Sense(&env); err != nil {
		println("error sensing")
		log.Fatal(err)
	}
	fmt.Printf("Sensor Temperature Reading: %8s \n", env.Temperature)

	resp := pb.BME280Response{Temperature: env.Temperature.String(), Pressure: env.Pressure.String(), Humidity: env.Humidity.String()}
	return &resp, nil
}

func (s *terragpioserver) SensePWM(ctx context.Context, pin *pb.PinSetRequest) (*pb.PWMResponse, error) {

	d := s.Pins[pin.PinNumber].DutyCycle.String()
	f := s.Pins[pin.PinNumber].Frequency.String()

	resp := pb.PWMResponse{Dutycycle: d, Frequency: f}
	return &resp, nil

}

// Set duty cycle on a pin based on the temperature reading from a BME280
func (s *terragpioserver) PWMDutyCycleOutput_BME280TempInput(ctx context.Context, settings *pb.FanControllerRequest) (*pb.FanControllerResponse, error) {
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
			r, err := s.SenseBME280(ctx, &pb.PinSetRequest{PinNumber: settings.BME280DevicePin})
			if err != nil {
				println("error sensing bme280")
				panic(err)
			}

			/* Temperature value will be returned as a string like "10 C"
			*  convert it to a physic.Temperature so we can convert it to a uint64 and do some math
			 */

			t.Set(r.Temperature)

			floatd := 0 - (float64(slope)*(float64(settings.TemperatureMax)-(t.Celsius())) - float64(settings.DutyCycleMax))

			intd := int(math.Round(floatd))
			stringd := strconv.Itoa(intd) + "%"

			d, err := gpio.ParseDuty(stringd)
			if err != nil {
				fmt.Println("error parsing duty cycle? : ", err)
				panic(err)
			}
			//d := settings.DutyCycleMax - (slope * (uint64(t.Celsius())))

			//set the dutycycle
			fan := s.Pins[settings.FanDevicePin]
			f.Set(fan.Frequency.String())
			s.setPWMDutyCycle(d,
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

	// Copied from https://pkg.go.dev/periph.io/x/conn/v3@v3.6.10/i2c/i2creg#All
	// Enumerate all I²C buses available and the corresponding pins.
	fmt.Print("I²C buses available:\n")
	for _, ref := range i2creg.All() {
		fmt.Printf("- %s\n", ref.Name)
		if ref.Number != -1 {
			fmt.Printf("  %d\n", ref.Number)
		}
		if len(ref.Aliases) != 0 {
			fmt.Printf("  %s\n", strings.Join(ref.Aliases, " "))
		}

		b, err := ref.Open()
		if err != nil {
			fmt.Printf("  Failed to open: %v", err)
		}
		if p, ok := b.(i2c.Pins); ok {
			fmt.Printf("  SDA: %s", p.SDA())
			fmt.Printf("  SCL: %s", p.SCL())
		}
		if err := b.Close(); err != nil {
			fmt.Printf("  Failed to close: %v", err)
		}
	}

	lis, err := net.Listen("tcp", "10.15.21.124:1234")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSetgpioServer(grpcServer, newServer())

	// Listen for client connections
	grpcServer.Serve(lis)
}
