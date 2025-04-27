package main

import (
	"github.com/aryanagn/ticket-go-backend/config"
	"github.com/aryanagn/ticket-go-backend/db"
	"github.com/aryanagn/ticket-go-backend/handlers"
	"github.com/aryanagn/ticket-go-backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator) // Replace db.DBMigrator with nil or the correct migrator function
	app := fiber.New(fiber.Config{
		AppName:      "TicketGoBook",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")

}
