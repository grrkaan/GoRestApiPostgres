package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/pkg/middleware"
	"github.com/grrkaan/go-rest-api/pkg/routes"
	"github.com/grrkaan/go-rest-api/platform/server"
)

func CreateServer(port int) {

	// Create Fiber App
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: server.ErrorHandler,
	})

	// Initialize Logger
	file := middleware.Logger(app)
	defer file.Close()

	routes.PostRoutes(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
