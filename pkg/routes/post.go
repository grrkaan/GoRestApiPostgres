package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/app/controllers/post"
)

func PostRoutes(app *fiber.App) {
	// List All Posts
	app.Get("/posts", post.FindAll)

	// Create Post
	app.Post("/post", post.Create)
}
