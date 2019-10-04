package admin

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	pageLimit = 5
)

func GetAdminArticles(c echo.Context) error {
	p, err := strconv.Atoi(c.QueryParam("page"));
	if err != nil || p == 0 {
		p = 1
	}

	ac := model.GetArticlesCount()
	tp := int(math.Ceil(float64(ac) / pageLimit))

	a := model.GetArticles((p - 1) * pageLimit, pageLimit)
	return c.Render(http.StatusOK, "page/admin/index", map[string]interface{}{
		"title": "Articles",
		"articles": a,
		"totalPage": tp,
		"currentPage": p,
	})
}

func GetAdminNewArticle(c echo.Context) error {
	ca := model.GetAllCategories()
	return c.Render(http.StatusOK, "page/admin/edit", map[string]interface{}{
		"title": "New article",
		"editable":   false,
		"categories": ca,
		"article":    model.ArticleFull{},
	})
}

func PostAdminNewArticle(c echo.Context) error {
	ca, err := strconv.Atoi(c.FormValue("category"))
	if err != nil {
		return err
	}

	s, err := sess.GetSession(c)
	if err != nil {
		return err
	}
	ui, ok := s.Values["user_id"]
	if !ok {
		return errors.New("invalid user id")
	}
	uis, ok := ui.(string)
	if !ok {
		return errors.New("invalid type")
	}
	u := model.GetUserByUserId(uis)

	err = model.InsertArticle(model.Article{
		Title:    c.FormValue("title"),
		Body:     c.FormValue("body"),
		CategoryID: uint(ca),
		UserID:   u.ID,
	})
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1452:") {
			c.Logger().Info("Invalid category id")
			return c.Redirect(http.StatusFound, "/admin/article")
		}
		return err
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func GetAdminArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ca := model.GetAllCategories()

	a, ok := model.GetArticle(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Article not found")
	}
	return c.Render(http.StatusOK, "page/admin/article", map[string]interface{}{
		"title": a.Title + " - SimpleBlog",
		"article":    a,
		"editable":   true,
		"categories": ca,
	})
}

func PostAdminArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	mode := c.FormValue("mode")
	if mode == "delete" {
		if err := model.DeleteArticle(id); err != nil {
			return err
		}
		return c.Redirect(http.StatusFound, "/admin/article")
	}

	ca, err := strconv.Atoi(c.FormValue("category"))
	if err != nil {
		return err
	}
	err = model.UpdateArticle(model.Article{
		Model: gorm.Model{
			ID: uint(id),
		},
		Title:    c.FormValue("title"),
		Body:     c.FormValue("body"),
		CategoryID: uint(ca),
	})
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1452:") {
			c.Logger().Info("Invalid category id")
			return c.Redirect(http.StatusFound, "/admin/article")
		}
		return err
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func DeleteAdminArticle(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteAdminArticle")
}
