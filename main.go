package main

import (
	"log"

	"github.com/crushr3sist/homely_backend/db"
	"github.com/crushr3sist/homely_backend/directory"
	"github.com/gofiber/fiber/v2"
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

	db.InitialMigration()
	app := appInstance()
	directory.DirRouter(app)
	log.Fatal(app.Listen(":8000"))
}
