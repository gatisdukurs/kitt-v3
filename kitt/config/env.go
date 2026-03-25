package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

const (
	ENV_PREFIX      = "kitt."
	ENV_NOT_SET     = ""
	ENV_DEVELOPMENT = "development"
	ENV_PRODUCTION  = "production"
	ENV_TESTING     = "testing"
)

func SetEnv(key string, value string) {
	os.Setenv(ENV_PREFIX+key, value)
}

func GetEnv(key string) string {
	return os.Getenv(ENV_PREFIX + key)
}

func LoadEnv(path string) error {
	path = strings.TrimSuffix(path, "/")

	env := os.Getenv(CONF_APP_ENV)
	if env == ENV_NOT_SET {
		env = ENV_DEVELOPMENT
		SetEnv(CONF_APP_ENV, env)
	}

	baseFile := filepath.Join(path, ".env")
	envFile := filepath.Join(path, ".env."+env)

	baseVars, err := godotenv.Read(baseFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	envVars, err := godotenv.Read(envFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	merged := map[string]string{}

	for k, v := range baseVars {
		merged[k] = v
	}

	for k, v := range envVars {
		merged[k] = v
	}

	for k, v := range merged {
		SetEnv(k, v)
	}

	return nil
}
