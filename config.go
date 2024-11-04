package log

import (
	"io"
	"os"

	"github.com/rafalb8/ln/util/env"
)

type Config struct {
	Level      *Level
	Format     string
	Enviroment string
	Multiline  *bool

	Output io.Writer
}

func (cfg *Config) defaults() {
	if cfg.Level == nil {
		lvl := LevelFrom(env.Get("LOG_LEVEL", "debug"))
		cfg.Level = &lvl
	}

	if cfg.Format == "" {
		cfg.Format = env.Get("LOG_FORMAT", "json")
	}

	if cfg.Enviroment == "" {
		cfg.Enviroment = env.Get("ENVIRONMENT", "dev")
	}

	if cfg.Multiline == nil {
		multiline := env.Get("LOG_MULTILINE", false)
		cfg.Multiline = &multiline
	}

	if cfg.Output == nil {
		cfg.Output = os.Stdout
	}
}
