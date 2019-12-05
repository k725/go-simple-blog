package main

import (
	"github.com/k725/go-simple-blog/config"
	"github.com/k725/go-simple-blog/controller/admin"
	"github.com/k725/go-simple-blog/controller/public"
	customMiddleware "github.com/k725/go-simple-blog/middleware"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/echoutil"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"log"
)

var dt = model.DbTarget{
	Address:  config.EnvDBAddress,
	User:     config.EnvDBUserName,
	Password: config.EnvDBPassword,
	Database: config.EnvDBName,
}

func main() {
	isDevelop := config.IsDevelopMode()
	if isDevelop {
		log.Println("Now development mode")
	}

	db, err := model.SetupConnection(dt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := model.CloseConnection(); err != nil {
			log.Fatal(err)
		}
	}()
	go model.SetupDB()

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB(), "session", "/", 3600, []byte(config.EnvSecret))
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	model.Settings = model.GetSettingValues()

	e := echo.New()
	e.Debug = isDevelop
	e.HideBanner = true
	e.HTTPErrorHandler = echoutil.CustomHTTPErrorHandler
	e.Renderer = echoutil.SetupPublicRender()

	e.Static("/", "public")
	e.Use(session.Middleware(store))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	registerRoutes(e)

	e.Logger.Fatal(e.Start(":8888"))
	defer func() {
		if err := e.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func registerRoutes(e *echo.Echo) {
	// Public
	e.GET("/", public.GetIndex)
	e.GET("/about", public.GetAbout)
	e.GET("/article/:id", public.GetArticle)
	e.GET("/category/:id", public.GetCategory)

	// Login
	e.GET("/admin/login", public.GetAdminLogin)
	e.POST("/admin/login", public.PostAdminLogin)

	// Login area
	g := e.Group(
		"/admin",
		customMiddleware.LoginCheck,
		echoutil.SetupAdminRender(),
	)

	g.GET("/logout", admin.GetAdminLogout)
	g.GET("/article", admin.GetAdminArticles)
	g.GET("/article/new", admin.GetAdminNewArticle)
	g.POST("/article/new", admin.PostAdminNewArticle)

	g.GET("/category", admin.GetCategories)
	g.POST("/category", admin.PostCategory)
	g.GET("/category/:id", admin.GetCategory)

	g.GET("/article/edit/:id", admin.GetAdminArticle)
	g.POST("/article/edit/:id", admin.PostAdminArticle)

	g.GET("/profile", admin.GetAdminProfile)
	g.POST("/profile", admin.PostAdminProfile)

	g.GET("/setting", admin.GetAdminSetting)
	g.POST("/setting", admin.PostAdminSetting)
}
