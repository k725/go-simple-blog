package util

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
)

func IsValidImageFormat(in multipart.File) bool {
	if _, _, err := image.DecodeConfig(in); err != nil {
		_, _ = in.Seek(0, io.SeekStart)
		return false
	}
	_, _ = in.Seek(0, io.SeekStart)
	return true
}