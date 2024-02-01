package handlers

import (
	"sort"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/gofiber/fiber/v2"
)

func HandleLanguages(c *fiber.Ctx) error {
	lexers := lexers.GlobalLexerRegistry.Lexers
	var languages []string

	for _, l := range lexers {
		config := l.Config()
		languages = append(languages, config.Aliases...)
	}
	sort.Sort(lexers)

	return c.Status(200).JSON(languages)
}
