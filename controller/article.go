package controller

import (
	"fmt"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	a := model.GetArticle(id)
	fmt.Println(a)
	return c.Render(http.StatusOK, "page/public/article", map[string]interface{}{
		"articles": a,
	})
}