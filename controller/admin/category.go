package admin

import (
	"fmt"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func GetCategories(c echo.Context) error {
	ca := model.GetAllCategories()

	s, err := sess.GetSession(c)
	if err != nil {
		return err
	}
	eF := s.Flashes("error")
	iF := s.Flashes("info")
	if err := sess.SaveSession(c, map[string]interface{}{}); err != nil {
		c.Error(err)
	}

	return echoview.Render(c, http.StatusOK, "page/admin/category", echo.Map{
		"title":      "Categories",
		"categories": ca,
		"errorFlash": eF,
		"infoFlash": iF,
	})
}

func PostCategory(c echo.Context) error {
	m := c.FormValue("mode")
	if m == "delete" {
		id := c.FormValue("id")
		idi, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		if err := model.DeleteCategory(idi); err != nil {
			if strings.HasPrefix(err.Error(), "Error 1451:") {
				c.Logger().Info("Category has articles.")
				if err := sess.SaveErrorFlash(c, "Category has articles."); err != nil {
					c.Logger().Warn(err)
				}
				return c.Redirect(http.StatusFound, "/admin/category")
			}
			return err
		}
		if err := sess.SaveInfoFlash(c, "Success delete."); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/category")
	}

	cn := c.FormValue("name")
	if len(cn) == 0 {
		if err := sess.SaveErrorFlash(c, "Invalid category name"); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/category")
	}

	ca := model.Category{
		Name: cn,
	}
	if err := model.InsertCategory(ca); err != nil {
		return err
	}
	if err := sess.SaveInfoFlash(c, fmt.Sprintf("Added category '%s'", cn)); err != nil {
		c.Logger().Warn(err)
	}
	return c.Redirect(http.StatusFound, "/admin/category")
}

func GetCategory(c echo.Context) error {
	p, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || p == 0 {
		p = 1
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if !model.HasCategory(id) {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	ac := model.GetArticlesByCategoryCount(id)
	tp := int(math.Ceil(float64(ac) / pageLimit))

	a := model.GetArticlesByCategory(id, (p-1)*pageLimit, pageLimit)
	return echoview.Render(c, http.StatusOK, "page/admin/index", echo.Map{
		"title":       "Category",
		"articles":    a,
		"totalPage":   tp,
		"currentPage": p,
	})
}
