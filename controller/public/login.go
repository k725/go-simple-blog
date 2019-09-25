package public

import (
	"errors"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/k725/go-simple-blog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminLogin(c echo.Context) error {
	s, err := sess.GetSession(c)
	if err != nil {
		return err
	}
	if _, ok := s.Values["user_id"]; ok {
		return c.Redirect(http.StatusFound, "/admin/article")
	}
	return c.Render(http.StatusOK, "page/public/login", map[string]interface{}{
		"title": "Login",
	})
}

func PostAdminLogin(c echo.Context) error {
	id := c.FormValue("user_id")
	pw := c.FormValue("password")
	if id == "" || pw == "" {
		return errors.New("username or password are empty")
	}

	user := model.GetUserByUserId(id)
	if err := util.PasswordVerify(user.Password, pw); err != nil {
		c.Logger().Warn(err)
		return c.Redirect(http.StatusFound, "/admin/login")
	}

	v := map[string]interface{}{
		"user_id": user.UserID,
	}
	if err := sess.SaveSession(c, v); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}
