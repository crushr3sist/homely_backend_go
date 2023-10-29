package main

import (
	"log"

	"github.com/crushr3sist/homely_backend/db"
	"github.com/crushr3sist/homely_backend/directory"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func appInstance() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "homely backend",
	})
	return app
}

func main() {

	app := appInstance()
	db.InitialMigration()
	// Add middleware
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Replace with your frontend URL
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Define the folder path for your video files
	videoFolderPath := "./videos" // Adjust the path as needed
	directory.DirRouter(app)
	// Serve static content (video files) using the fiber.Static middleware
	app.Static("/streams", videoFolderPath)

	log.Fatal(app.Listen(":8000"))
}
