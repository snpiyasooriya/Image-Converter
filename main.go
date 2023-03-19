package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	imageconvert "github.com/snpiyasooriya/webp-converter/image-convert"
)

func main() {
	app := fiber.New()

	app.Post("/convert-image", func(c *fiber.Ctx) error {
		c.Accepts("json", "text")     // "json"
		c.Accepts("application/json") // "application/json"

		payload := struct {
			ImageUrl string `json:"imageUrl"`
			Function string `json:"function"`
		}{}

		c.BodyParser(&payload)
		var webp imageconvert.Webp
		webp.WebpConvert(payload.ImageUrl, payload.Function)
		fmt.Println(&payload.ImageUrl)
		return c.JSON(webp)
		// return nil

	})

	app.Listen(":3000")

}
