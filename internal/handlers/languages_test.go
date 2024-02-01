package handlers

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_LanguagesEndPointWorks(t *testing.T) {
	app := fiber.New()
	app.Get("/languages", HandleLanguages)

	request := httptest.NewRequest("GET", "/languages", nil)

	response, err := app.Test(request, 100)

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	assert.NotNil(t, body)
}
