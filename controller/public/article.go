package public

import (
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

	a := model.GetArticle(id)
	ca := model.GetAllCategories()

	a.Body = markdown.Render(a.Body)
	return c.Render(http.StatusOK, "page/public/article", map[string]interface{}{
		"article": a,
		"categories": ca,
	})
}
