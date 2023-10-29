package directory

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/crushr3sist/homely_backend/convert"
	"github.com/gofiber/fiber/v2"
)

func parseDir(c *fiber.Ctx) error {

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
	convert.ConvertAll(parsedDirs.RawPaths)
	fmt.Print(parsedDirs)
	return c.JSON(parsedDirs)

}

type Show struct {
	Name     string   `json:"name"`
	Episodes []string `json:"episodes"`
}

type ShowResponse struct {
	Shows []Show `json:"shows"`
}

func ReturnShowsList() ShowResponse {
	showsDir := "./videos"

	showsList, err := ioutil.ReadDir(showsDir)
	if err != nil {
		log.Fatalf("Error reading shows directory: %v", err)
	}

	var shows []Show

	for _, showInfo := range showsList {
		if showInfo.IsDir() {
			showName := showInfo.Name()
			episodesDir := filepath.Join(showsDir, showName)

			episodesList, err := ioutil.ReadDir(episodesDir)
			if err != nil {
				log.Fatalf("Error reading episodes directory for show %s: %v", showName, err)
			}

			var episodes []string
			for _, episodeInfo := range episodesList {
				episodes = append(episodes, episodeInfo.Name())
			}

			shows = append(shows, Show{Name: showName, Episodes: episodes})
		}
	}

	return ShowResponse{Shows: shows}
}

func returnDirs(c *fiber.Ctx) error {
	data := ReturnShowsList()
	return c.Status(fiber.StatusOK).JSON(data)
}

func DirRouter(app *fiber.App) {
	dirAPI := app.Group("/dir")

	dirAPI.Get("/parse", parseDir)
	dirAPI.Get("/all_videos", returnDirs)

}
