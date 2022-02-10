package config

import "os"

const (
	envKey = "ENV"
	// EnvDevelopment is the development environment
	EnvDevelopment = "local"
	// EnvProduction is the production environment
	EnvProduction = "production"
)

var env = GetEnv(envKey, EnvDevelopment)

func Env() string {
	return env
}

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)

	if ok {
		return env
	}

	return def
}
