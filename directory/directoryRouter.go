package directory

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func returnDirs(c *fiber.Ctx) error {
	// gets the path from request : done
	// walks the dir if found : done
	// we save the path into the database
	//
	c.Accepts("application/json")

	dataStruct := struct {
		ContentType       string `json:"contentType"`
		DirectoryToTarget string `json:"directoryToTarget"`
	}{}

	if err := c.BodyParser(&dataStruct); err != nil {
		fmt.Println("error :", err)
		return c.SendStatus(500)
	}

	parsedDirs := ParseDirs(dataStruct.DirectoryToTarget, dataStruct.ContentType)

	fmt.Print(parsedDirs)

	return c.JSON(parsedDirs)

}

func DirRouter(app *fiber.App) {
	dirAPI := app.Group("/dir")

	dirAPI.Get("/parse", returnDirs)
}
