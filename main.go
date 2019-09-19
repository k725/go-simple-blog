package main

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/config"
	"github.com/k725/go-simple-blog/controller/admin"
	"github.com/k725/go-simple-blog/controller/public"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"html/template"
	"net/http"
	"time"
	"unicode/utf8"
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

	// model.SetupDB()

	e := echo.New()

	con := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", t.User, t.Password, t.Address, t.Database)
	store, err := mysqlstore.NewMySQLStore(con, "session", "/", 3600, []byte("secret"))
	if err != nil {
		panic(err)
	}

	defer store.Close()

	e.Use(session.Middleware(store))

	e.Debug = true
	e.Static("/", "public")

	setupRender(e)
	setupMiddleware(e)
	setupRoute(e)

	e.Logger.Fatal(e.Start(":8888"))
	// defer e.Close()
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
			"copy": func() string {
				return time.Now().Format("2006")
			},
			"dateYYYYMMDD": func(t time.Time) string {
				return t.Format("2006/01/02")
			},
			"dateYYYYMMDDHHmm": func(t time.Time) string {
				return t.Format("2006/01/02 15:04")
			},
			"eqTime": func(t1, t2 time.Time) bool {
				return t1.Equal(t2)
			},
			"trimChars": func(t1 string, len int) string {
				if utf8.RuneCountInString(t1) <= len {
					return t1
				}
				return string([]rune(t1)[0:len])
			},
			"safeHTML": func(t string) template.HTML {
				return template.HTML(t)
			},
		},
		DisableCache: true,
	})
}

func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
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
	g := e.Group("/admin")
	g.Use(ServerHeader)
	g.GET("/logout", admin.GetAdminLogout)
	g.GET("/article", admin.GetAdminArticles)
	g.GET("/article/new", admin.GetAdminNewArticle)
	g.POST("/article/new", admin.PostAdminNewArticle)

	g.GET("/article/edit/:id", admin.GetAdminArticle)
	g.POST("/article/edit/:id", admin.PostAdminArticle)
	g.DELETE("/article/edit/:id", admin.DeleteAdminArticle)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s, err := sess.GetSession(c)
		if err != nil {
			// @todo
			return err
		}
		if _, ok := s.Values["user_id"]; !ok {
			return c.Redirect(http.StatusFound, "/admin/login")
		}

		err = next(c)
		// After
		return err
	}
}
