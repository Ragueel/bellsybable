package highlighter

import (
	"bytes"
	"errors"
	"io"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var ErrInvalidLanguage = errors.New("invalid language")

type GenerateRequest struct {
	Code        string
	Language    string
	Style       string
	WithClasses bool
}

type GenerateResult struct {
	Content string
	Styles  string
}

func NewDefaultRequest(code, language, style string) GenerateRequest {
	return GenerateRequest{
		Code:        code,
		Language:    language,
		Style:       style,
		WithClasses: false,
	}
}

func GenerateHtmlHighlight(request GenerateRequest) (*GenerateResult, error) {
	lexer := lexers.Get(request.Language)

	if lexer == nil {
		return nil, ErrInvalidLanguage
	}

	style := styles.Get(request.Style)

	codeBuffer := bytes.NewBufferString(request.Code)
	contents, err := io.ReadAll(codeBuffer)
	if err != nil {
		return nil, err
	}

	iterator, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		return nil, err
	}

	var contentWriter bytes.Buffer
	formatter := html.New(html.WithClasses(request.WithClasses))
	err = formatter.Format(&contentWriter, style, iterator)
	if err != nil {
		return nil, err
	}

	var stylesWriter bytes.Buffer
	err = formatter.WriteCSS(&stylesWriter, style)
	if err != nil {
		return nil, err
	}

	result := GenerateResult{
		Content: contentWriter.String(),
		Styles:  stylesWriter.String(),
	}
	return &result, nil
}
