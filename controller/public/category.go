package public

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

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

	ca := model.GetCategoryById(id)


	ac := model.GetArticlesByCategoryCount(id)
	tp := int(math.Ceil(float64(ac) / pageLimit))

	a := model.GetArticlesByCategory(id, (p-1)*pageLimit, pageLimit)

	// @todo temp template
	return echoview.Render(c, http.StatusOK, "page/public/index", echo.Map{
		"title":        ca.Name,
		"articles":     a,
		"categories":   model.GetAllCategories(),
		"totalPage":    tp,
		"currentPage":  p,
		"categoryName": ca.Name,
	})
}
