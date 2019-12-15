package public

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/grokify/html-strip-tags-go"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/markdown"
	"github.com/k725/go-simple-blog/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
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

	desc := []rune(strip.StripTags(a.Body))
	if len(desc) >= 160 {
		desc = desc[:160]
	}
	ogpDesc :=  strings.Replace(string(desc), "\n", "", -1)

	return echoview.Render(c, http.StatusOK, "page/public/article", echo.Map{
		"title":      a.Title,
		"article":    a,
		"categories": ca,
		"ogp":        map[string]interface{} {
			"title": a.Title,
			"type": "article",
			"url": "https://example.com/article/" + strconv.Itoa(id),
			"thumbnail": "https://example.com/article.png",
			"site_name": util.GetSettingValue("blog-name", ""),
			"description": ogpDesc,
		},
	})
}
