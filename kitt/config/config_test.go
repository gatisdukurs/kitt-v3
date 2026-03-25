package config

import (
	"os"
	"testing"
)

func Test_Config(t *testing.T) {
	t.Run("it sets and gets", func(t *testing.T) {
		cfg := NewConfig()
		cfg.Set("foo", "bar")

		have := cfg.Get("foo")
		want := "bar"

		if have != want {
			t.Fatalf("!eq: %s -> %s", have, want)
		}

		if cfg.Has("does_not_exist") != false {
			t.Fatal("should be false")
		}
	})

	t.Run("it loads env variables", func(t *testing.T) {
		LoadEnv("./testdata")
		conf := NewConfigFromEnv()

		have := conf.Get(CONF_APP_ENV)
		want := "development"

		if have != want {
			t.Fatalf("!eq: %s -> %s", have, want)
		}

		os.Setenv(CONF_APP_ENV, ENV_PRODUCTION)
		LoadEnv("./testdata")
		conf = NewConfigFromEnv()

		have = conf[CONF_APP_HOST]
		want = "prod.app"

		if have != want {
			t.Fatalf("!eq: %s -> %s", have, want)
		}
	})
}
