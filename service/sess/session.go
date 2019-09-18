package sess

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetSession(c echo.Context) (*sessions.Session, error) {
	s, err := session.Get("session", c)
	if err != nil {
		return nil, err
	}
	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
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
