package admin

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/model"
	"github.com/k725/go-simple-blog/service/sess"
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
	user := model.GetUserByUserId(uis)
	return c.Render(http.StatusOK, "page/admin/profile", map[string]interface{}{
		"user": user,
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
			ID:        uinf.ID,
		},
		UserID:   c.FormValue("user-id"),
		Name:     c.FormValue("display-name"),
	}
	if err := model.UpdateUser(u); err != nil {
		return err;
	}
	return c.Redirect(http.StatusFound, "/admin/profile")
}
