package util

import "strings"

func TrimSpace(s string) string {
	return strings.Trim(s, "\r\n\t ")
}
