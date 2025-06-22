package ln

import (
	"io"
	"os"

	"github.com/rafalb8/ln/util/env"
)

type Config struct {
	Level       *Level
	Format      string
	Environment string
	Multiline   *bool

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

	if cfg.Environment == "" {
		cfg.Environment = env.Get("ENVIRONMENT", "dev")
	}

	if cfg.Multiline == nil {
		multiline := env.Get("LOG_MULTILINE", false)
		cfg.Multiline = &multiline
	}

	if cfg.Output == nil {
		cfg.Output = os.Stdout
	}
}
