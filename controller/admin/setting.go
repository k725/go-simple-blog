package admin

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/k725/go-simple-blog/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminSetting(c echo.Context) error {
	m := echo.Map{}

	for _, v := range model.GetSettingValues() {
		m[v.Key] = v.Value
	}

	return echoview.Render(c, http.StatusOK, "page/admin/setting", m)
}

func PostAdminSetting(c echo.Context) error {
	f, _ := c.FormParams()
	allowValues := []string {
		"blog-name",
		"blog-description",
		"google-analytics",
		"site-url",
	}
	for _, v := range allowValues {
		setting := model.Setting{
			Key:   v,
			Value: f.Get(v),
		}
		_ = model.UpdateSettingValue(setting)
	}
	model.Settings = model.GetSettingValues()

	return c.Redirect(http.StatusFound, "/admin/setting")
}

