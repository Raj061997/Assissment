package app

import (
	"example/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	Init()
	con := controller.NewController(application.service)
	api := app.Group("/api")

	api.Post("/blog-post", con.CreatePost)
	api.Get("/blog-post", con.GetPosts)
	api.Get("/blog-post/:id", con.GetPost)
	api.Patch("/blog-post/:id", con.UpdatePost)
	api.Delete("/blog-post/:id", con.DeletePost)
}
