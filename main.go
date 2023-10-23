package main

import (
	"log"
	"path"

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

	var newPath string = path.Base("g:/adventure time/")

	directory.FormatPath(newPath)

	db.InitialMigration()

	app := appInstance()

	log.Fatal(app.Listen(":3000"))
}
