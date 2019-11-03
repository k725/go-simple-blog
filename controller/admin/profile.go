package admin

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
	"github.com/k725/go-simple-blog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminProfile(c echo.Context) error {
	id, err := sess.GetSessionValue(c, "user_id")
	if err != nil {
		c.Logger().Warn(err)
		return err
	}
	user := model.GetUserByUserId(id)
	return echoview.Render(c, http.StatusOK, "page/admin/profile", echo.Map{
		"user": user,
		"errorFlash": sess.GetFlash(c, "error"),
		"infoFlash":  sess.GetFlash(c, "info"),
	})
}

func PostAdminProfile(c echo.Context) error {
	id, err := sess.GetSessionValue(c, "user_id")
	if err != nil {
		c.Logger().Warn(err)
		return c.Redirect(http.StatusFound, "/admin/profile")
	}

	uinf := model.GetUserByUserId(id)
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
