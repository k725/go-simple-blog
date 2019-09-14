package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/login", map[string]interface{}{})
}

func PostAdminLogin(c echo.Context) error {
	uv, err := c.FormParams()
	if err != nil {
		return err
	}
	// @todo Unimplemented
	if uv.Get("user_id") != "admin" || uv.Get("password") != "dolphin" {
		return c.String(http.StatusForbidden, "Missing auth")
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func GetAdminArticles(c echo.Context) error {
	return c.String(http.StatusOK, "GetAdminArticles")
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
