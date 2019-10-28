package admin

import (
	"errors"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/k725/go-simple-blog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminProfile(c echo.Context) error {
	s, err := sess.GetSession(c)
	if err != nil {
		return err
	}
	ui, ok := s.Values["user_id"]
	if !ok {
		return errors.New("invalid user id")
	}
	uis, ok := ui.(string)
	if !ok {
		return errors.New("invalid type")
	}

	eF := s.Flashes("error")
	iF := s.Flashes("info")
	if err := sess.SaveSession(c, map[string]interface{}{}); err != nil {
		c.Error(err)
	}

	user := model.GetUserByUserId(uis)
	return echoview.Render(c, http.StatusOK, "page/admin/profile", echo.Map{
		"user": user,
		"errorFlash": eF,
		"infoFlash": iF,
	})
}

func PostAdminProfile(c echo.Context) error {
	s, err := sess.GetSession(c)
	if err != nil {
		return err
	}

	ui, ok := s.Values["user_id"]
	if !ok {
		return errors.New("invalid user id")
	}

	uis, ok := ui.(string)
	if !ok {
		return errors.New("invalid type")
	}

	uinf := model.GetUserByUserId(uis)

	u := model.User{
		Model: gorm.Model{
			ID: uinf.ID,
		},
		UserID: c.FormValue("user-id"),
		Name:   c.FormValue("display-name"),
	}

	if c.FormValue("old-password") != "" && c.FormValue("new-password") != "" {
		if err := util.PasswordVerify(uinf.Password, c.FormValue("old-password")); err != nil {
			c.Logger().Warn(err)
			if err := sess.SaveErrorFlash(c, "Missing old password"); err != nil {
				c.Logger().Warn(err)
			}
			return c.Redirect(http.StatusFound, "/admin/profile")
		}

		pw, err := util.PasswordHash(c.FormValue("new-password"))
		if err != nil {
			c.Logger().Warn(err)
			if err := sess.SaveErrorFlash(c, "Failed hash new password"); err != nil {
				c.Logger().Warn(err)
			}
			return c.Redirect(http.StatusFound, "/admin/profile")
		}
		u.Password = pw
	}

	if err := model.UpdateUser(u); err != nil {
		c.Logger().Warn(err)
		if err := sess.SaveErrorFlash(c, "Failed update user info."); err != nil {
			c.Logger().Warn(err)
		}
		return c.Redirect(http.StatusFound, "/admin/profile")
	}
	if err := sess.SaveInfoFlash(c, "Success update."); err != nil {
		c.Logger().Warn(err)
	}
	if c.FormValue("old-password") != "" && c.FormValue("new-password") != "" {
		if err := sess.SaveInfoFlash(c, "Password update."); err != nil {
			c.Logger().Warn(err)
		}
	}
	return c.Redirect(http.StatusFound, "/admin/profile")
}
