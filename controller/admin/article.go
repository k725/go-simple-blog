package admin

import (
	"github.com/foolin/goview/supports/echoview-v4"
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
	p, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || p == 0 {
		p = 1
	}

	ac := model.GetArticlesCount()
	tp := int(math.Ceil(float64(ac) / pageLimit))

	a := model.GetArticles((p-1)*pageLimit, pageLimit)
	return echoview.Render(c, http.StatusOK, "page/admin/index", echo.Map{
		"title":       "Articles",
		"articles":    a,
		"totalPage":   tp,
		"currentPage": p,
		"errorFlash":  sess.GetFlash(c, "error"),
		"infoFlash":   sess.GetFlash(c, "info"),
	})
}

func GetAdminNewArticle(c echo.Context) error {
	return echoview.Render(c, http.StatusOK, "page/admin/edit", echo.Map{
		"title":      "New article",
		"editable":   false,
		"categories": model.GetAllCategories(),
		"article":    model.ArticleFull{},
	})
}

func PostAdminNewArticle(c echo.Context) error {
	ca, err := strconv.Atoi(c.FormValue("category"))
	if err != nil {
		c.Logger().Warn(err)
		if err := sess.SaveErrorFlash(c, "Invalid format category id"); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/article")
	}

	id, err := sess.GetSessionValue(c, "user_id")
	if err != nil {
		c.Logger().Warn(err)
		return c.Redirect(http.StatusFound, "/admin/article")
	}
	u := model.GetUserByUserId(id)

	err = model.InsertArticle(model.Article{
		Title:      c.FormValue("title"),
		Body:       c.FormValue("body"),
		TopImage:   c.FormValue("top-image"),
		CategoryID: uint(ca),
		UserID:     u.ID,
	})
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1452:") {
			if err := sess.SaveErrorFlash(c, "Invalid category id"); err != nil {
				c.Logger().Warn(err)
			}
			c.Logger().Info("Invalid category id")
			return c.Redirect(http.StatusFound, "/admin/article")
		}
		return err
	}
	if err := sess.SaveInfoFlash(c, "Created article"); err != nil {
		c.Logger().Warn(err)
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}

func GetAdminArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Warn(err)
		return echo.NewHTTPError(http.StatusNotFound, "Invalid article ID")
	}

	a, ok := model.GetArticle(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Article not found")
	}
	return echoview.Render(c, http.StatusOK, "page/admin/article", echo.Map{
		"title":      a.Title + " - SimpleBlog",
		"article":    a,
		"editable":   true,
		"categories": model.GetAllCategories(),
	})
}

func PostAdminArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Warn(err)
		return echo.NewHTTPError(http.StatusNotFound, "Invalid article ID")
	}

	mode := c.FormValue("mode")
	if mode == "delete" {
		if err := model.DeleteArticle(id); err != nil {
			c.Logger().Warn(err)
			if err := sess.SaveErrorFlash(c, "Failed delete article"); err != nil {
				c.Logger().Warn(err)
			}
			return c.Redirect(http.StatusFound, "/admin/article")
		}

		if err := sess.SaveInfoFlash(c, "Successful delete article"); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/article")
	}

	ca, err := strconv.Atoi(c.FormValue("category"))
	if err != nil {
		c.Logger().Warn(err)
		if err := sess.SaveErrorFlash(c, "Invalid format category id"); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/article")
	}
	err = model.UpdateArticle(model.Article{
		Model: gorm.Model{
			ID: uint(id),
		},
		Title:      c.FormValue("title"),
		Body:       c.FormValue("body"),
		TopImage:   c.FormValue("top-image"),
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
