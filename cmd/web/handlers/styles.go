package handlers

import (
	"github.com/alecthomas/chroma/styles"
	"github.com/gofiber/fiber/v2"
)

func HandleStyles(c *fiber.Ctx) error {
	allStyles := styles.Names()

	return c.Status(200).JSON(allStyles)
}
