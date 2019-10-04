package main

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/config"
	"github.com/k725/go-simple-blog/controller/admin"
	"github.com/k725/go-simple-blog/controller/public"
	customMiddleware "github.com/k725/go-simple-blog/middleware"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/util/templateutil"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"log"
	"net/http"
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
	db.LogMode(isDevelop)
	db.AutoMigrate(
		&model.Article{},
		&model.User{},
		&model.Category{},
	)
	db.Model(&model.Article{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Article{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	go model.SetupDB()

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB(), "session", "/", 3600, []byte(config.EnvSecret))
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	e := echo.New()
	e.Debug = isDevelop
	e.HideBanner = true
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Renderer = setupRender()

	e.Static("/", "public")
	e.Use(session.Middleware(store))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setupRoute(e)

	e.Logger.Fatal(e.Start(":8888"))
	defer func () {
		if err := e.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func setupRender() *echoview.ViewEngine{
	return echoview.New(goview.Config{
		Root:      "template",
		Extension: ".html",
		Master:    "base",
		Partials:  []string{
			// "partials/ad",
		},
		Funcs: templateutil.TemplateFuncMap,
		DisableCache: true,
	})
}

func setupRoute(e *echo.Echo) {
	// Public
	e.GET("/", public.GetIndex)
	e.GET("/about", public.GetAbout)
	e.GET("/article/:id", public.GetArticle)
	e.GET("/category/:id", public.GetCategory)

	// Login
	e.GET("/admin/login", public.GetAdminLogin)
	e.POST("/admin/login", public.PostAdminLogin)

	// Login area
	g := e.Group("/admin", customMiddleware.LoginCheck)

	g.GET("/logout", admin.GetAdminLogout)
	g.GET("/article", admin.GetAdminArticles)
	g.GET("/article/new", admin.GetAdminNewArticle)
	g.POST("/article/new", admin.PostAdminNewArticle)

	g.GET("/category", admin.GetCategories)
	g.POST("/category", admin.PostCategory)
	g.GET("/category/:id", admin.GetCategory)

	g.GET("/article/edit/:id", admin.GetAdminArticle)
	g.POST("/article/edit/:id", admin.PostAdminArticle)
	g.DELETE("/article/edit/:id", admin.DeleteAdminArticle)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	// Error page
	errorPage := fmt.Sprintf("page/error/%d", code)
	errorParam := map[string]interface{}{
		"title": code,
	}
	if err := c.Render(code, errorPage, errorParam); err != nil {
		c.Logger().Error(err)
	} else {
		return
	}

	// エラーページのテンプレートレンダリングで失敗したときの保険
	if err := c.String(code, "Internal Server Error"); err != nil {
		c.Logger().Error(err)
	}
}
