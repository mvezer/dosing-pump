package pump

import (
	"encoding/json"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/api"
	"gobot.io/x/gobot/v2/drivers/i2c"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

type RunPumpParams struct {
	Id       int           `json:"id"`
	Duration time.Duration `json:"duration"`
}

type Pump struct {
	master      *gobot.Master
	bot         *gobot.Robot
	RpiAdaptor  *raspi.Adaptor
	MotorDriver *i2c.Adafruit2348Driver
	Speed       int32
}

func extractParam(params map[string]interface{}, target any) error {
	jsonBytes, mError := json.Marshal(params)
	if mError != nil {
		return mError
	}

	uError := json.Unmarshal(jsonBytes, target)
	if uError != nil {
		return uError
	}

	return nil
}

func (p *Pump) RunPump(motorId int, duration time.Duration) {
	p.MotorDriver.SetDCMotorSpeed(int(motorId), p.Speed)
	p.MotorDriver.RunDCMotor(int(motorId), i2c.Adafruit2348Forward)
	if duration > 0 {
		time.Sleep(duration * time.Millisecond)
		p.MotorDriver.RunDCMotor(int(motorId), i2c.Adafruit2348Release)
	}
}

func NewPump() *Pump {
	pump := Pump{}
	pump.master = gobot.NewMaster()
	api.NewAPI(pump.master).Start()
	pump.RpiAdaptor = raspi.NewAdaptor()
	pump.MotorDriver = i2c.NewAdafruit2348Driver(pump.RpiAdaptor)
	pump.Speed = 128

	pump.bot = pump.master.AddRobot(gobot.NewRobot("pump",
		[]gobot.Connection{pump.RpiAdaptor},
		[]gobot.Device{pump.MotorDriver},
	))
	pump.bot.AddCommand("run_pump", func(params map[string]interface{}) interface{} {
		p := RunPumpParams{}
		err := extractParam(params, &p)

		if err != nil {
			return err
		}

		if p.Id < 0 || p.Id > 3 {
			return "ERROR: the id must be between [0-3]"
		}

		if p.Duration < 10 || p.Duration > 10000 {
			return "ERROR: the duration must in between 10 and 10000 ms"
		}

		pump.RunPump(p.Id, p.Duration)

		return "OK"
	})

	pump.bot.Start()

	return &pump
}

func (p *Pump) GetMaster() *gobot.Master {
	return p.master
}
