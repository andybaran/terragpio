package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

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

func (s *terragpioserver) PWMDutyCycleOutput_BME280TempInput(ctx context.Context, settings *pb.FanControllerRequest) (*pb.FanControllerResponse, error) {
	//setup the PWM device
	s.SetPWM(ctx, settings.FanDevice)

	//setup the BME280
	// ToDo : Should have SetBME280 and a SenseBME280...not all in one
	r, err := s.GetBME280(ctx, settings.BME280Device)
	if err != nil {
		panic(err)
	}

	//calculate our slope
	slope := (settings.TemperatureMax - settings.DutyCycleMax) / (settings.TemperatureMin - settings.DutyCycleMin)
	println("s = ", slope)

	/* We want to start a loop here that gets the temp and sets the duty cycle
	*  However, we don't want to be in a blocking loop so the loop can be brought into a go routine
	 */

	//calculate duty cycle (y axis using y = mx+b)
	var t physic.Temperature
	t.Set(r.Temperature)
	d := (slope*(uint64(t.Celsius())-settings.TemperatureMax) + settings.DutyCycleMax)

	//set the dutycycle

	var f physic.Frequency
	f.Set(settings.FanDevice.Frequency)
	setPWMDutyCycle(gpio.Duty(d),
		f,
		gpioreg.ByName(settings.FanDevice.Pin))

	resp := pb.FanControllerResponse{}
	return &resp, nil
}

/* func (s *terragpioserver) SetFanController(ctx context.Context, settings *pb.FanControllerRequest) (*pb.FanControllerResponse, error) {

//calculate our slope
slope := (int(settings.TemperatureMax) - int(settings.DutyCycleMax)) / (int(settings.TemperatureMin) - int(settings.DutyCycleMin))
println("s = ", slope)

/* We want to start a loop here that gets the temp and sets the duty cycle
*  However, we don't want to be in a blocking loop so the loop can be brought into a go routine
*/

/*d := (slope*(int(c.Celsius())-tMax) + dMax)
	//calculate duty cycle (y axis using y = mx+b)
	setPWMDutyCycle(gpio.Duty(d),
		25000,
		gpioreg.ByName("GPIO13"))

} */

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

/* func readBME() (physic.Temperature, error) {
	bus, err := i2creg.Open("") //just open the first bus found
	//fmt.Println("i2c bus opened")
	if err != nil {
		return 0, err
	}
	defer bus.Close()

	dev, err := bmxx80.NewI2C(bus, uint16(0x77), &bmxx80.DefaultOpts) //0x77 is default for the bme280 I currently have
	//fmt.Println("ready to read values")
	if err != nil {
		return 0, err
	}
	defer dev.Halt()

	// Read temperature from the sensor:
	var env physic.Env
	if err = dev.Sense(&env); err != nil {
		return 0, err
	}
	//	fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)
	//	fmt.Println("returning env")
	return env.Temperature, nil
}
*/
func setPWMDutyCycle(d gpio.Duty, f physic.Frequency, p gpio.PinIO) error {

	if err := p.PWM(d, f); err != nil {
		println(err)
		return err
	}
	println("duty cycle: ", d)
	println()
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
			//fmt.Println("Recieved temperature of: ", t)
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
			s := (tMax - dMax) / (tMin - dMin)
			println("s = ", s)

			d := (s*(int(c.Celsius())-tMax) + dMax)
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

	myTicker := time.NewTicker(time.Second * 2)
	for range myTicker.C {
		actualTemp, err := readBME()
		if err != nil {
			panic(err)
		}
		temperatureChan <- actualTemp
		println("temp = ", actualTemp.String())
	}
	*/
}
