package main

import (
	"github.com/mvezer/dosing-pump/internal/pump"
)

func main() {
	// master := gobot.NewMaster()
	// server.NewServer(master)
	pump.NewPump()

	// fmt.Printf("Pump instance robots count: %d\n", pump.GetMaster().Robots().Len())
	// fmt.Printf("Api port: %s", api.Port)

}
