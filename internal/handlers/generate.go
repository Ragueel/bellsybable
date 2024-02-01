package handlers

import (
	"bellsybabble/internal/highlighter"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type GenerateRequest struct {
	Code     string `validate:"required" json:"code"`
	Language string `validate:"required" json:"language"`
	Style    string `validate:"required" json:"style"`
}

type GenerateResponse struct {
	Code string `json:"code"`
}

type GenerateHandler struct {
	validator *validator.Validate
}

type ErrorResonse struct {
	Message string `json:"message"`
}

func NewGenerateHandler(validator *validator.Validate) *GenerateHandler {
	return &GenerateHandler{
		validator: validator,
	}
}

func (handler *GenerateHandler) Handle(c *fiber.Ctx) error {
	request := GenerateRequest{}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(ErrorResonse{Message: "Invalid body"})
	}

	if err := handler.validator.Struct(request); err != nil {
		return c.Status(400).JSON(ErrorResonse{Message: fmt.Sprintf("validation failed: %s", err)})
	}

	settings := highlighter.NewDefaultSettings(request.Code, request.Language, request.Style)

	result, err := highlighter.GenerateHtmlHighlight(settings)
	if err != nil {
		return c.Status(500).JSON(ErrorResonse{Message: fmt.Sprintf("error creating code: %s", err)})
	}

	return c.Status(200).JSON(GenerateResponse{Code: result.Content})
}
