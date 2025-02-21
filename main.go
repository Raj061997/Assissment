package main

import (
	"example/database"
	_ "example/docs" // Import the generated docs
	"example/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // Import Fiber Swagger
	"github.com/joho/godotenv"
	"log"
)

// @title Blog CRUD API
// @version 1.0
// @description Simple Blog API using Go-Fiber, PostgreSQL, and Swagger
// @host localhost:3000
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	database.ConnectDatabase()
	router.SetupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins (or specify frontend URL)
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	// ✅ Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault) // This serves Swagger UI
	//	log.Printf("Starting server on port %s...", 10000)
	app.Listen(":10000")
}
