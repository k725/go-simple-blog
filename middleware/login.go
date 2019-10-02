package middleware

import (
	"github.com/gorilla/securecookie"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginCheck(next echo.HandlerFunc) echo.HandlerFunc {
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
