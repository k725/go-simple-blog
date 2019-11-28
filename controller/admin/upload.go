package admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PostUploadFile(c echo.Context) error {
	return c.JSON(http.StatusOK, &echo.Map{
		"status": true,
	})
}