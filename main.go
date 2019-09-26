package main

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/gorilla/securecookie"
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
	"log"
	"net/http"
	"time"
	"unicode/utf8"
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
		panic(err)
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

	// model.SetupDB()

	e := echo.New()

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB(), "session", "/", 3600, []byte(config.EnvSecret))
	if err != nil {
		panic(err)
	}
	defer store.Close()

	e.Use(session.Middleware(store))

	e.Debug = isDevelop
	e.Static("/", "public")

	e.HTTPErrorHandler = customHTTPErrorHandler

	setupRender(e)
	setupMiddleware(e)
	setupRoute(e)

	e.Logger.Fatal(e.Start(":8888"))
	defer func () {
		if err := e.Close(); err != nil {
			log.Fatal(err)
		}
	}()
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
			"dateToLocal": func(t time.Time) time.Time {
				l, _ := time.LoadLocation("Local")
				return t.In(l)
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
			"add": func(v1, v2 int) int {
				return v1 + v2
			},
			"sub": func(v1, v2 int) int {
				return v1 - v2
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
	g := e.Group("/admin", loginCheckMiddleware)

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

func loginCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s, err := sess.GetSession(c)
		if err != nil && err.Error() == securecookie.ErrMacInvalid.Error() {
			if err := sess.ForceLogoutSession(c); err != nil {
				return err
			}
			return c.Redirect(http.StatusFound, "/admin/login")
		} else if err != nil {
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

