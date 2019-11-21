package admin

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAdminSetting(c echo.Context) error {
	return echoview.Render(c, http.StatusOK, "page/admin/setting", echo.Map{})
}

