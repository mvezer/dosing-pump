package server

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
)

func NewServer(m *gobot.Master) *api.API {
	server := api.NewAPI(m)
	// server.Port = "3000"
	// server.AddHandler(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
	// })
	// server.Debug()
	server.Start()

	return server
}
