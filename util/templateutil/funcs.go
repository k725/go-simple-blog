package templateutil

import (
	"html/template"
	"time"
	"unicode/utf8"
)

func getCurrentYear() string {
	return time.Now().Format("2006")
}

func dateToLocal(t time.Time) time.Time {
	l, _ := time.LoadLocation("Local")
	return t.In(l)
}

func dateYYYYMMDD(t time.Time) string {
	return t.Format("2006/01/02")
}

func dateYYYYMMDDHHmm(t time.Time) string {
	return t.Format("2006/01/02 15:04")
}

func equalDate(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

func trimChars(t1 string, len int) string {
	if utf8.RuneCountInString(t1) <= len {
		return t1
	}
	return string([]rune(t1)[0:len])
}

func safeHTML(t string) template.HTML {
	return template.HTML(t)
}

func add(v1, v2 int) int {
	return v1 + v2
}

func sub(v1, v2 int) int {
	return v1 - v2
}