package public

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/markdown"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	a, ok := model.GetArticle(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Article not found")
	}

	ca := model.GetAllCategories()
	a.Body = markdown.Render(a.Body)
	return echoview.Render(c, http.StatusOK, "page/public/article", echo.Map{
		"title":      a.Title + " - SimpleBlog",
		"article":    a,
		"categories": ca,
	})
}
