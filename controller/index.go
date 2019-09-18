package controller

import (
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIndex(c echo.Context) error {
	a := model.GetAllArticles()
	return c.Render(http.StatusOK, "page/public/index", map[string]interface{}{
		"articles": a,
	})
}

func GetAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/about", map[string]interface{}{
		"title": "This About title",
	})
}
