package handlers

import (
	"bellsybabble/internal/highlighter"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type GenerateRequest struct {
	Code     string `validate:"required"`
	Language string `validate:"required"`
	Style    string `validate:"required"`
}

type GenerateResponse struct {
	Code string `json:"code"`
}

type GenerateHandler struct {
	xValidator *validator.Validate
}

func (handler *GenerateHandler) Handle(c *fiber.Ctx) error {
	request := GenerateRequest{}

	if err := c.BodyParser(request); err != nil {
		return err
	}

	err := handler.xValidator.Struct(request)
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
