package templateutil

import (
	"html/template"
)

var TemplateFuncMap = template.FuncMap{
	"copy": getCurrentYear,
	"dateToLocal": dateToLocal,
	"dateYYYYMMDD": dateYYYYMMDD,
	"dateYYYYMMDDHHmm": dateYYYYMMDDHHmm,
	"eqTime": equalDate,
	"trimChars": trimChars,
	"safeHTML": safeHTML,
	"add": add,
	"sub": sub,
}
