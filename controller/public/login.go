package public

import (
	"errors"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/gorilla/securecookie"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/k725/go-simple-blog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminLogin(c echo.Context) error {
	p := map[string]interface{}{
		"title": "Login",
	}

	s, err := sess.GetSession(c)
	if err == nil {
		if _, ok := s.Values["user_id"]; ok {
			return c.Redirect(http.StatusFound, "/admin/article")
		}
		p["errorFlash"] = sess.GetFlash(c, "error")
		p["infoFlash"] = sess.GetFlash(c, "info")
		return echoview.Render(c, http.StatusOK, "page/public/login", p)
	}

	if err.Error() == securecookie.ErrMacInvalid.Error() {
		if err := sess.ForceLogoutSession(c); err != nil {
			return err
		}

		// NOTE: c.Render を最後に呼び出さないとCookieが書き換わらないので😭
		return echoview.Render(c, http.StatusOK, "page/public/login", p)
	}
	return err
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
		if err := sess.SaveErrorFlash(c, "Missing username or password"); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/login")
	}

	if err := sess.SaveInfoFlash(c, "Successful login"); err != nil {
		c.Logger().Warn(err)
	}

	v := map[string]interface{}{
		"user_id": user.UserID,
	}
	if err := sess.SaveSession(c, v); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/admin/article")
}
