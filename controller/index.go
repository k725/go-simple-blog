package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/index", map[string]interface{}{
		"title": "This Index title",
		"body": "This Body text",
	})
}

func GetAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/about", map[string]interface{}{
		"title": "This About title",
	})
}

func GetTest(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}