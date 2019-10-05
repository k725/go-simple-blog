package echoutil

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	// Error page
	errorPage := fmt.Sprintf("page/error/%d", code)
	errorParam := map[string]interface{}{
		"title": code,
	}
	if err := c.Render(code, errorPage, errorParam); err != nil {
		c.Logger().Error(err)
	} else {
		return
	}

	// エラーページのテンプレートレンダリングで失敗したときの保険
	if err := c.String(code, "Internal Server Error"); err != nil {
		c.Logger().Error(err)
	}
}
