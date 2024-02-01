package handlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateHandlerProperlyWorks(t *testing.T) {
	app := fiber.New()

	validator := validator.New()
	handler := NewGenerateHandler(validator)

	app.Post("/generate", handler.Handle)

	reqBody := GenerateRequest{
		Code:     `echo "hello world"`,
		Language: "bash",
		Style:    "monokai",
	}
	result, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	sendData := string(result)
	request := httptest.NewRequest("POST", "/generate", strings.NewReader(sendData))
	request.Header.Add("Content-Length", strconv.FormatInt(request.ContentLength, 10))
	request.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(request, 100)

	assert.NoError(t, err)
	assert.NotNil(t, resp)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)

	assert.NotNil(t, body)
}
