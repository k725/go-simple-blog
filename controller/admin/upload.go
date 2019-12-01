package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/k725/go-simple-blog/util"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func PostUploadFile(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if !util.IsValidImageFormat(src) {
		return c.JSON(http.StatusOK, &echo.Map{
			"status": false,
		})
	}

	hashName := fmt.Sprintf("%x", md5.Sum([]byte(file.Filename)))
	dstPath := filepath.Join("public", "image", "article", hashName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"status": true,
		"data": echo.Map{
			"filePath":	"image/article/" + hashName,
		},
	})
}