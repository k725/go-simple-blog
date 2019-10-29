package sess

import (
	"github.com/gorilla/sessions"
	"github.com/k725/go-simple-blog/config"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const cookieName = "session"

func GetSession(c echo.Context) (*sessions.Session, error) {
	s, err := session.Get(cookieName, c)
	if err != nil {
		return nil, err
	}
	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		Secure:   !config.IsDevelopMode(),
		HttpOnly: true,
	}
	return s, nil
}

func SaveSession(c echo.Context, d map[string]interface{}) error {
	sess, err := GetSession(c)
	if err != nil {
		return err
	}
	for k, v := range d {
		sess.Values[k] = v
	}
	return sess.Save(c.Request(), c.Response())
}

func SaveErrorFlash(c echo.Context, d string) error {
	return saveFlash(c, d, "error")
}

func SaveInfoFlash(c echo.Context, d string) error {
	return saveFlash(c, d, "info")
}

func saveFlash(c echo.Context, d string, key ...string) error {
	sess, err := GetSession(c)
	if err != nil {
		return err
	}
	sess.AddFlash(d, key...)
	return sess.Save(c.Request(), c.Response())
}

func GetFlash(c echo.Context, key string) []interface{} {
	s, err := GetSession(c)
	if err != nil {
		c.Logger().Error(err)
		return []interface{}{}
	}
	f := s.Flashes(key)
	if err := SaveSession(c, echo.Map{}); err != nil {
		c.Logger().Error(err)
		return []interface{}{}
	}
	return f
}

func ForceLogoutSession(c echo.Context) error {
	s, err := session.Get(cookieName, c)
	if err != nil {
		c.Logger().Warn(err)
	}
	return s.Save(c.Request(), c.Response())
}
