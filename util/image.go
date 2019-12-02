package util

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
)

func IsValidImageFormat(in multipart.File) (string, bool) {
	_, ext, err := image.DecodeConfig(in)
	if err != nil {
		_, _ = in.Seek(0, io.SeekStart)
		return "", false
	}
	_, _ = in.Seek(0, io.SeekStart)
	return ext, true
}