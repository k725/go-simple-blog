package markdown

import (
	"github.com/russross/blackfriday/v2"
	"strings"
)

func Render(text string) string {
	text = strings.Replace(text, "\r\n", "\n", -1)
	html := blackfriday.Run([]byte(text))
	return string(html)
}
