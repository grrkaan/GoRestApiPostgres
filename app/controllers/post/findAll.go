package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/app/models"
	"github.com/grrkaan/go-rest-api/platform/database"
)

func FindAll(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	database.Connection().Find(&posts)

	return ctx.JSON(posts)
}
