package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/pkg/utils"
	"github.com/grrkaan/go-rest-api/platform/migrations"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if !fiber.IsChild() {
		// Application initialization
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
