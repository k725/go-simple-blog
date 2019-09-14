package controller

import (
	"fmt"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIndex(c echo.Context) error {
	var a []model.Article
	con := model.GetConnection()
	con.Find(&a)

	for _, v := range a {
		fmt.Println(v.Title, v.Body, v.Author, v.Category, v.CreatedAt.String())
	}

	return c.Render(http.StatusOK, "page/public/index", map[string]interface{}{
		"title": "This Index title",
		"body": "This Body text",
	})
}

func GetAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/about", map[string]interface{}{
		"title": "This About title",
	})
}

func GetTest(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}