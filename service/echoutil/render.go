package echoutil

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/util/templateutil"
	"github.com/labstack/echo/v4"
)

func SetupPublicRender() *echoview.ViewEngine {
	return echoview.New(goview.Config{
		Root:      "template",
		Extension: ".html",
		Master:    "base_public",
		Partials:  []string{
			// "partials/ad",
		},
		Funcs:        templateutil.TemplateFuncMap,
		DisableCache: true,
	})
}

func SetupAdminRender() echo.MiddlewareFunc {
	return echoview.NewMiddleware(goview.Config{
		Root:      "template",
		Extension: ".html",
		Master:    "base_admin",
		Partials:  []string{
			// "partials/ad",
		},
		Funcs:        templateutil.TemplateFuncMap,
		DisableCache: true,
	})
}
