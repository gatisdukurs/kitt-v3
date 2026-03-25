package config

import (
	"os"
	"testing"
)

func Test_Env(t *testing.T) {
	t.Run("it loads env variables", func(t *testing.T) {
		os.Setenv(CONF_APP_ENV, "")

		if err := LoadEnv("./testdata"); err != nil {
			t.Fatal(err)
		}

		have := GetEnv(CONF_APP_ENV)
		want := "development"

		if have != want {
			t.Fatalf("!eq: %s -> %s", have, want)
		}

		have = GetEnv(CONF_APP_HOST)
		want = "dev.app"

		if have != want {
			t.Fatalf("!eq: %s -> %s", have, want)
		}
	})
}
