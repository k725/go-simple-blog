package public

import (
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIndex(c echo.Context) error {
	a := model.GetAllArticles()
	ca := model.GetAllCategories()
	return c.Render(http.StatusOK, "page/public/index", map[string]interface{}{
		"articles": a,
		"categories": ca,
	})
}

func GetAbout(c echo.Context) error {
	ca := model.GetAllCategories()
	return c.Render(http.StatusOK, "page/public/about", map[string]interface{}{
		"title": "This About title",
		"categories": ca,
	})
}
