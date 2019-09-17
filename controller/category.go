package controller

import (
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	a := model.GetArticlesByCategory(id)
	// @todo temp template
	return c.Render(http.StatusOK, "page/public/index", map[string]interface{}{
		"articles": a,
	})
}
