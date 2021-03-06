package admin

import (
	"crypto/md5"
	"errors"
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

	if file.Size > 2145728 { // @todo temp
		return errors.New("file too large")
	}

	ext, ok := util.IsValidImageFormat(src)
	if !ok {
		return c.JSON(http.StatusOK, &echo.Map{
			"status": false,
		})
	}

	hashName := fmt.Sprintf("%x.%s", md5.Sum([]byte(file.Filename)), ext)
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
			"filePath":	"/image/article/" + hashName,
		},
	})
}