package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/aryanagn/ticket-go-app/backend/config"
	"github.com/aryanagn/ticket-go-app/backend/db"
	"github.com/aryanagn/ticket-go-app/backend/handlers"
	"github.com/aryanagn/ticket-go-app/backend/middlewares"
	"github.com/aryanagn/ticket-go-app/backend/repositories"
	"github.com/aryanagn/ticket-go-app/backend/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
