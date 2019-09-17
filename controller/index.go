package controller

import (
	"github.com/gorilla/sessions"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIndex(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	sess.Save(c.Request(), c.Response())

	a := model.GetAllArticles()
	return c.Render(http.StatusOK, "page/public/index", map[string]interface{}{
		"articles": a,
	})
}

func GetAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "page/public/about", map[string]interface{}{
		"title": "This About title",
	})
}
