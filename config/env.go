package config

import "os"

var (
	EnvDBAddress  string
	EnvDBUserName string
	EnvDBPassword string
	EnvDBName     string
)

func init() {
	EnvDBAddress = mustGetEnv("DB_ADDRESS")
	EnvDBUserName = mustGetEnv("DB_USERNAME")
	EnvDBPassword = mustGetEnv("DB_PASSWORD")
	EnvDBName = mustGetEnv("DB_NAME")
}

func mustGetEnv(k string) string {
	r := os.Getenv(k)
	if len(r) <= 0 {
		panic("Environment variable '" + k + "' is not set.")
	}
	return r
}
