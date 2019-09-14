package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetArticle(c echo.Context) error {
	return c.String(http.StatusOK, "GetArticle")
}