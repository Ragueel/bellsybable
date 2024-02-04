package handlers

import "github.com/gofiber/fiber/v2"

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HandleHealtchCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(HealthCheckResponse{Message: "ok"})
}
