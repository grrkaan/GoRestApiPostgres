package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/pkg/routes"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New()

	routes.PostRoutes(app)
	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
