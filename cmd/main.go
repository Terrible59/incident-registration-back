package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"incidents_back/pkg/handlers"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Incidents",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		IdleTimeout: time.Minute,
	})

	handlers.SetupRoutes(&handlers.SetupConfig{
		App: app,
	})

	if err := app.Listen(":8888"); err != nil {
		log.WithError(err).Fatal("Server failed to start!")
	}
}
