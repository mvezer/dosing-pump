package db

import (
	"fmt"

	"gobot.io/x/gobot/v2/drivers/i2c"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type Pump struct {
	Id int
	// Button       *gpio.ButtonDriver
	Motor        *i2c.AdafruitMotorHatDriver
	ButtonEvents chan interface{}
}

type PumpController struct {
	Pumps       map[int]Pump
	RpiAdaptor  *raspi.Adaptor
	MotorSpeed  int32
	MotorDriver *i2c.AdafruitMotorHatDriver
}

func Init() PumpController {
	pc := PumpController{
		Pumps:       make(map[int]Pump),
		RpiAdaptor:  raspi.NewAdaptor(),
		MotorSpeed:  128,
		MotorDriver: i2c.NewAdafruitMotorHatDriver(raspi.NewAdaptor()),
	}
	pc.MotorDriver.Start()

	return pc
}

func (pc *PumpController) Start() {
	pc.MotorDriver.Start()
}

func (pc *PumpController) AddPump(id int, buttonPin string) {
	pc.Pumps[id] = Pump{
		Id: id,
		// Button: gpio.NewButtonDriver(pc.RpiAdaptor, buttonPin),
	}
	// pc.Pumps[id].Button.Start()
	// pc.Pumps[id].Button.On(gpio.ButtonPush, func(data interface{}) {
	// 	fmt.Printf("Pump %d button pushed\n", id)
	// 	pc.RunPump(id, 0)
	// })
	// pc.Pumps[id].Button.On(gpio.ButtonRelease, func(data interface{}) {
	// 	fmt.Printf("Pump %d button released\n", id)
	// 	pc.StopPump(id)
	// })
	pc.Pumps[id].Motor.Start()
}

func (pc *PumpController) RunPump(id int, interval int) {
	fmt.Printf("Running pump %d\n", id)
	if err := pc.MotorDriver.SetDCMotorSpeed(id, pc.MotorSpeed); err != nil {
		fmt.Printf("Error setting motor speed: %s\n", err)
	}
	if err := pc.MotorDriver.RunDCMotor(id, i2c.AdafruitForward); err != nil {
		fmt.Printf("Error running motor: %s\n", err)
	}
}

func (pc *PumpController) StopPump(id int) {
	fmt.Printf("Stopping pump %d\n", id)
	if err := pc.MotorDriver.RunDCMotor(id, i2c.AdafruitRelease); err != nil {
		fmt.Printf("Error stopping the motor: %s\n", err)
	}
}

// func adafruitDCMotorRunner(a *i2c.AdafruitMotorHatDriver, dcMotor int) (err error) {
// 	log.Printf("DC Motor Run Loop...\n")
// 	// set the speed:
// 	var speed int32 = 255 // 255 = full speed!
// 	if err = a.SetDCMotorSpeed(dcMotor, speed); err != nil {
// 		return
// 	}
// 	// run FORWARD
// 	if err = a.RunDCMotor(dcMotor, i2c.AdafruitForward); err != nil {
// 		return
// 	}
// 	// Sleep and RELEASE
// 	time.Sleep(2000 * time.Millisecond)
// 	if err = a.RunDCMotor(dcMotor, i2c.AdafruitRelease); err != nil {
// 		return
// 	}
// 	// run BACKWARD
// 	if err = a.RunDCMotor(dcMotor, i2c.AdafruitBackward); err != nil {
// 		return
// 	}
// 	// Sleep and RELEASE
// 	time.Sleep(2000 * time.Millisecond)
// 	if err = a.RunDCMotor(dcMotor, i2c.AdafruitRelease); err != nil {
// 		return
// 	}
// 	return
// }
//
// func InitMotorController() {
// 	r := raspi.NewAdaptor()
// 	adaFruit := i2c.NewAdafruitMotorHatDriver(r)
//
// 	work := func() {
// 		gobot.Every(5*time.Second, func() {
//
// 			dcMotor := 2 // 0-based
// 			adafruitDCMotorRunner(adaFruit, dcMotor)
// 		})
// 	}
//
// 	robot := gobot.NewRobot("adaFruitBot",
// 		[]gobot.Connection{r},
// 		[]gobot.Device{adaFruit},
// 		work,
// 	)
//
// 	robot.Start()
// }
