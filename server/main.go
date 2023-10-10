package main

import (
	// "strconv"

	"github.com/gofiber/fiber/v2"

	config "github.com/mvezer/dosing-pump/config"
	// pumpcontroller "github.com/mvezer/dosing-pump/pumpcontroller"
	"gopkg.in/yaml.v2"
)

// var PumpController = pumpcontroller.Init()
var CurrentConfig *config.Config

func setupRoutes(app *fiber.App) {
	app.Post("/api/config", setConfig)
	app.Get("/api/config", getConfig)
	// app.Post("/api/pump/:id", runPump)
}

func main() {
	// engine := html.New("./templates", ".html")
	app := fiber.New()
	// app.Static("/", "../client/public")
	// app.Get("/", mainPage)
	setupRoutes(app)

	app.Listen(":5000")
}

func setConfig(c *fiber.Ctx) error {
	configYaml := c.Body()
	CurrentConfig, err := config.ParseConfig(configYaml)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	config.SaveConfig(CurrentConfig)

	return c.JSON("ok")
}

func getConfig(c *fiber.Ctx) error {
	yamlData, err := yaml.Marshal(CurrentConfig)
	if err != nil {
		c.Status(503).SendString(err.Error())
	}
	return c.SendString(string(yamlData))
}

// func runPump(c *fiber.Ctx) error {
// 	id, cerr := strconv.ParseInt(c.Params("id"), 10, 12)
// 	if cerr != nil {
// 		return c.Status(400).SendString(cerr.Error())
// 	}
//
// 	PumpController.RunPump(int(id), 10)
//
// 	return c.JSON("ok")
// }
