package server

import (
	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/api"
)

func NewServer(m *gobot.Master) *api.API {
	server := api.NewAPI(m)
	server.Port = "3000"
	server.Debug()
	server.Start()

	return server
}
