package public

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

const (
	pageLimit = 5
)

func GetIndex(c echo.Context) error {
	p, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || p == 0 {
		p = 1
	}

	ac := model.GetArticlesCount()
	tp := int(math.Ceil(float64(ac) / pageLimit))

	a := model.GetArticles((p-1)*pageLimit, pageLimit)
	ca := model.GetAllCategories()
	return echoview.Render(c, http.StatusOK, "page/public/index", echo.Map{
		"title":       "SimpleBlog",
		"articles":    a,
		"categories":  ca,
		"totalPage":   tp,
		"currentPage": p,
	})
}

func GetAbout(c echo.Context) error {
	ca := model.GetAllCategories()
	return echoview.Render(c, http.StatusOK, "page/public/about", echo.Map{
		"title":      "This About title",
		"categories": ca,
	})
}
