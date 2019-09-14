package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/config"
	"github.com/k725/go-simple-blog/controller"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"time"
)

func main() {

	t := model.DbTarget{
		Address:  config.EnvDBAddress,
		User:     config.EnvDBUserName,
		Password: config.EnvDBPassword,
		Database: config.EnvDBName,
	}
	db, err := model.SetupConnection(t)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := model.CloseConnection(); err != nil {
			panic(err)
		}
	}()
	db.LogMode(true)
	db.AutoMigrate(
		&model.Article{},
		&model.User{},
		&model.Category{},
	)

	//test := model.Article{
	//	Title:    "日記",
	//	Body:     "今日は晴れてた🌞",
	//	Category: 0,
	//	Author:   0,
	//}
	//db.NewRecord(test)
	//db.Create(&test)


	e := echo.New()

	e.Debug = true
	e.Static("/", "public")

	setupRender(e)
	setupMiddleware(e)
	setupRoute(e)

	e.Logger.Fatal(e.Start(":8888"))
}

func setupRender(e *echo.Echo) {
	e.Renderer = echoview.New(goview.Config{
		Root:      "template",
		Extension: ".html",
		Master:    "base",
		Partials:  []string{
			// "partials/ad",
		},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
			"sushi": func(a string) string {
				return "🍣" + a + "🍣"
			},
		},
		DisableCache: true,
	})
}

func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func setupRoute(e *echo.Echo)  {
	e.GET("/", controller.GetIndex)
	e.GET("/about", controller.GetAbout)
	e.GET("/test", controller.GetTest)
}