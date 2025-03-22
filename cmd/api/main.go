package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kwangminnam/rn-with-go/config"
	"github.com/kwangminnam/rn-with-go/db"
	"github.com/kwangminnam/rn-with-go/handlers"
	"github.com/kwangminnam/rn-with-go/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrate)
	app := fiber.New(fiber.Config{
		AppName:      "RN with Go",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/events"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
