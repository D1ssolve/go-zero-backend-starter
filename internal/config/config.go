package config

import (
	"os"
	"strings"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Postgres struct {
		DSN string
	}
	Telegram struct {
		Token string `json:",optional"`
	}
}

func (c *Config) ExpandEnvironment() {
	c.Postgres.DSN = os.ExpandEnv(c.Postgres.DSN)
	c.Telegram.Token = os.ExpandEnv(c.Telegram.Token)
	if value := strings.TrimSpace(os.Getenv("DATABASE_URL")); value != "" {
		c.Postgres.DSN = value
	}
	if value := strings.TrimSpace(os.Getenv("TELEGRAM_BOT_TOKEN")); value != "" {
		c.Telegram.Token = value
	}
}
