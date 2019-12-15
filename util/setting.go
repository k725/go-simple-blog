package util

import "github.com/k725/go-simple-blog/model"

func GetSettingValue(key, def string) string {
	for _, v := range model.Settings {
		if v.Key == key {
			return v.Value
		}
	}
	return def
}
