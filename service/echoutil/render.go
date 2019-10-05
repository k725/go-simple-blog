package echoutil

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/util/templateutil"
)

func SetupRender() *echoview.ViewEngine {
	return echoview.New(goview.Config{
		Root:      "template",
		Extension: ".html",
		Master:    "base",
		Partials:  []string{
			// "partials/ad",
		},
		Funcs:        templateutil.TemplateFuncMap,
		DisableCache: true,
	})
}
