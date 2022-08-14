package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/app/models"
	"github.com/grrkaan/go-rest-api/platform/database"
)

func FindById(ctx *fiber.Ctx) error {

	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error

	if err != nil {
		return err
	}

	return ctx.JSON(post)
}
