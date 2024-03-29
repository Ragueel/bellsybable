package handlers

import (
	"sort"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/gofiber/fiber/v2"
)

func HandleLanguages(c *fiber.Ctx) error {
	lexers := lexers.GlobalLexerRegistry.Lexers
	var languages []string

	sort.Sort(lexers)

	for _, l := range lexers {
		config := l.Config()
		languages = append(languages, config.Aliases...)
	}

	return c.Status(200).JSON(languages)
}
