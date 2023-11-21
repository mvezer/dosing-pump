package pump

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/api"
	"gobot.io/x/gobot/v2/drivers/i2c"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type Pump struct {
	master      *gobot.Master
	bot         *gobot.Robot
	RpiAdaptor  *raspi.Adaptor
	MotorDriver *i2c.Adafruit2348Driver
	Speed       int32
}

func NewPump() *Pump {
	pump := Pump{}
	pump.master = gobot.NewMaster()
	api.NewAPI(pump.master).Start()
	pump.RpiAdaptor = raspi.NewAdaptor()
	pump.MotorDriver = i2c.NewAdafruit2348Driver(pump.RpiAdaptor)
	pump.Speed = 128

	pump.master.AddCommand("custom_command", func(params map[string]interface{}) interface{} {

		return "This command is attached to the mcp!"
	})

	pump.bot = pump.master.AddRobot(gobot.NewRobot("pump",
		[]gobot.Connection{pump.RpiAdaptor},
		[]gobot.Device{pump.MotorDriver},
	))
	pump.bot.AddCommand("run_pump", func(params map[string]interface{}) interface{} {
		fmt.Println("Alright...")

		motorIdInterface, ok := params["id"]
		if !ok {
			return "ERROR: you need to specify a pump id!"
		}

		motorIdFloat, ok := motorIdInterface.(float64)
		if !ok {
			return "ERROR: wrong id type!"
		}

		motorId := int(motorIdFloat)

		if int(motorId) < 0 || int(motorId) > 3 {
			return "ERROR: the id must be between [0-3]"
		}
		fmt.Println("Alright...")

		pump.MotorDriver.SetDCMotorSpeed(int(motorId), pump.Speed)
		pump.MotorDriver.RunDCMotor(int(motorId), i2c.Adafruit2348Forward)
		time.Sleep(2000 * time.Millisecond)
		pump.MotorDriver.RunDCMotor(int(motorId), i2c.Adafruit2348Release)
		return "OK"
	})

	pump.bot.Start()
	pump.master.Start()

	return &pump
}

func (p *Pump) GetMaster() *gobot.Master {
	return p.master
}
