package config

import (
	"os"
	"strings"
)

const (
	CONF_APP_ENV  = "app.env"
	CONF_APP_NAME = "app.name"
	CONF_APP_HOST = "app.host"
)

type Config map[string]string

func NewConfig() Config {
	return Config{}
}

func (c Config) Set(key string, value string) Config {
	c[key] = value
	return c
}

func (c Config) Get(key string) string {
	if c.Has(key) {
		return c[key]
	}
	return ""
}

func (c Config) Has(key string) bool {
	_, ok := c[key]
	return ok
}

func NewConfigFromEnv() Config {
	cfg := NewConfig()
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		}

		if !strings.HasPrefix(parts[0], ENV_PREFIX) {
			continue
		}

		key := strings.TrimPrefix(parts[0], ENV_PREFIX)
		value := parts[1]

		// Transform key: APP_ENV → app.env
		k := strings.ToLower(key)
		k = strings.ReplaceAll(k, "_", ".")

		cfg.Set(k, value)
	}

	return cfg
}
