package controller

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/k725/go-simple-blog/util"
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

	id := uv.Get("user_id")
	pw := uv.Get("password")
	if id == "" || pw == "" {
		return errors.New("username or password are empty")
	}
	user := model.GetUserByUserId(id)

	if err := util.PasswordVerify(user.Password, pw); err != nil {
		c.Logger().Error(err)
		return c.Redirect(http.StatusFound, "/admin")
	}

	v := map[string]interface{}{
		"login": "ok",
		"true": true,
		"false": false,
		"nil": nil,
		"int": 123,
	}
	if err := sess.SaveSession(c, v); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func GetAdminArticles(c echo.Context) error {
	// s, _ := sess.GetSession(c)

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
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
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
		Model: gorm.Model{
			ID: uint(id),
		},
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	})
	if err != nil {
		return nil
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func DeleteAdminArticle(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteAdminArticle")
}
