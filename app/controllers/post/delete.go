package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/app/models"
	"github.com/grrkaan/go-rest-api/platform/database"
)

func Delete(ctx *fiber.Ctx) error {

	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error

	if err != nil {
		return err
	}

	err = database.Connection().Delete(&post).Error

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
