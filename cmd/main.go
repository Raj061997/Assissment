package main

import (
	engin "example/cmd/app"
	_ "example/docs" // Import the generated docs
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger" // Import Fiber Swagger
	"github.com/joho/godotenv"
)

// @title Blog CRUD API
// @version 1.0
// @description Simple Blog API using Go-Fiber, PostgreSQL, and Swagger
// @host assissment-xpx7.onrender.com
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins (or specify frontend URL)
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" // Default port if not set
	}
	// ✅ Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault) // This serves Swagger UI
	engin.SetupRoutes(app)

	log.Printf("Starting server on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
