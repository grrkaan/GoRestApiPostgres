package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/app/models"
	"github.com/grrkaan/go-rest-api/platform/database"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}

	if err := ctx.BodyParser(&post); err != nil {
		return ctx.Status(503).JSON(err)
	}

	if err := database.Connection().Create(&post).Error; err != nil {
		return ctx.Status(503).JSON(err)
	}

	return ctx.JSON(post)
}
