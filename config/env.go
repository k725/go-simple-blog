package config

import (
	"github.com/k725/go-simple-blog/util"
	"os"
)

var (
	EnvDBAddress  string
	EnvDBUserName string
	EnvDBPassword string
	EnvDBName     string
	EnvSecret     string
	Env           string
)

func init() {
	EnvDBAddress = mustGetEnv("DB_ADDRESS")
	EnvDBUserName = mustGetEnv("DB_USERNAME")
	EnvDBPassword = mustGetEnv("DB_PASSWORD")
	EnvDBName = mustGetEnv("DB_NAME")
	EnvSecret = defaultGetEnv("SECRET", "uMo@ASV#x6!UaWF")
	Env = defaultGetEnv("ENV", development)
}

func mustGetEnv(k string) string {
	r := os.Getenv(k)
	if len(r) <= 0 {
		panic("Environment variable '" + k + "' is not set.")
	}
	return util.TrimSpace(r)
}

func defaultGetEnv(k, def string) string {
	r := os.Getenv(k)
	if len(r) <= 0 {
		return def
	}
	return util.TrimSpace(r)
}
