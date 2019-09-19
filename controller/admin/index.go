package admin

import (
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminLogout(c echo.Context) error {
	s, err := sess.GetSession(c)
	if err != nil {
		return err
	}
	s.Options.MaxAge = -1
	if err := s.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/admin/login")
}
