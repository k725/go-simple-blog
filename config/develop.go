package config

const development = "development"

func IsDevelopMode() bool {
	return Env == development
}
