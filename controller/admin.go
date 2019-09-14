package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/login", map[string]interface{}{})
}

func PostAdminLogin(c echo.Context) error {
	return c.String(http.StatusOK, "PostAdminLogin")
}

func GetAdminArticle(c echo.Context) error {
	return c.String(http.StatusOK, "GetAdminArticle")
}

func PostAdminArticle(c echo.Context) error {
	return c.String(http.StatusOK, "PostAdminArticle")
}

func DeleteAdminArticle(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteAdminArticle")
}
