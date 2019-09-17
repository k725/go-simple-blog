package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	a := model.GetAllArticles()
	return c.Render(http.StatusOK, "page/admin/index", map[string]interface{}{
		"articles": a,
	})
}

func GetAdminNewArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "page/admin/edit", map[string]interface{}{})
}

func PostAdminNewArticle(c echo.Context) error {
	err := model.InsertArticle(model.Article{
		Title:    c.FormValue("title"),
		Body:     c.FormValue("body"),
	})
	if err != nil {
		return nil
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func GetAdminArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	a := model.GetArticle(id)
	return c.Render(http.StatusOK, "page/admin/article", map[string]interface{}{
		"article": a,
	})
}

func PostAdminArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = model.UpdateArticle(model.Article{
		Model:    gorm.Model{
			ID:        uint(id),
		},
		Title:    c.FormValue("title"),
		Body:     c.FormValue("body"),
	})
	if err != nil {
		return nil
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func DeleteAdminArticle(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteAdminArticle")
}
