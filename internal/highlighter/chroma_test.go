package highlighter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HighlighterWorks(t *testing.T) {
	code := "echo \"Hello world\""
	language := "bash"
	style := "rose-pine-moon"
	expectedContent := `<pre tabindex="0" style="color:#e5e5e5;background-color:#000;"><code><span style="display:flex;"><span><span style="color:#fff;font-weight:bold">echo</span> <span style="color:#0ff;font-weight:bold">&#34;Hello world&#34;</span></span></span></code></pre>`

	result, err := GenerateHtmlHighlight(code, language, style)

	assert.NoError(t, err)

	fmt.Println(result.Content)
	fmt.Println(result.Styles)

	assert.Equal(t, expectedContent, result.Content)
	assert.True(t, false)
}
