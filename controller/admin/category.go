package admin

import (
	"errors"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

func GetCategories(c echo.Context) error {
	ca := model.GetAllCategories()
	return c.Render(http.StatusOK, "page/admin/category", map[string]interface{}{
		"title": "Categories",
		"categories": ca,
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

		if !model.HasCategory(idi) {

		}
		if err := model.DeleteCategory(idi); err != nil {
			return err
		}
		return c.Redirect(http.StatusFound, "/admin/category")
	}

	cn := c.FormValue("name")
	if len(cn) == 0 {
		return errors.New("invalid category name")
	}

	ca := model.Category{
		Name:  cn,
	}
	if err := model.InsertCategory(ca); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/admin/category")
}

func GetCategory(c echo.Context) error {
	p, err := strconv.Atoi(c.QueryParam("page"));
	if err != nil || p == 0 {
		p = 1
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	ac := model.GetArticlesByCategoryCount(id)
	tp := int(math.Ceil(float64(ac) / pageLimit))

	a := model.GetArticlesByCategory(id, (p - 1) * pageLimit, pageLimit)
	return c.Render(http.StatusOK, "page/admin/index", map[string]interface{}{
		"title": "Category",
		"articles": a,
		"totalPage": tp,
		"currentPage": p,
	})
}