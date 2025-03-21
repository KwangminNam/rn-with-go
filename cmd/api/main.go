package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kwangminnam/rn-with-go/handlers"
	"github.com/kwangminnam/rn-with-go/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "RN with Go",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/events"), eventRepository)

	app.Listen(":3000")
}
