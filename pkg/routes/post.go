package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grrkaan/go-rest-api/app/controllers/post"
)

func PostRoutes(app *fiber.App) {
	// List All Posts
	app.Get("/posts", post.FindAll)

	// Create Post
	app.Post("/post/create", post.Create)

	// Update Post
	app.Put("/post/update/:id", post.Update)

	// Find By Id Post
	app.Get("/post/:id", post.FindById)

	// Delete Post
	app.Delete("/post/delete/:id", post.Delete)
}
