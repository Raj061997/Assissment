package router

import (
	"example/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/blog-post", controller.CreatePost)
	api.Get("/blog-post", controller.GetPosts)
	api.Get("/blog-post/:id", controller.GetPost)
	api.Patch("/blog-post/:id", controller.UpdatePost)
	api.Delete("/blog-post/:id", controller.DeletePost)
}
