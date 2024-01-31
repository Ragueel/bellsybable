package handlers

import (
	"bellsybabble/internal/highlighter"

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

func NewGenerateHandler(validator *validator.Validate) *GenerateHandler {
	return &GenerateHandler{
		validator: validator,
	}
}

func (handler *GenerateHandler) Handle(c *fiber.Ctx) error {
	request := GenerateRequest{}

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	err := handler.validator.Struct(request)
	if err != nil {
		return err
	}

	settings := highlighter.NewDefaultSettings(request.Code, request.Language, request.Style)

	result, err := highlighter.GenerateHtmlHighlight(settings)
	if err != nil {
		return err
	}

	return c.JSON(GenerateResponse{Code: result.Content})
}
