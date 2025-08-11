package ln

import (
	"io"
	"os"

	"github.com/rafalb8/ln/util/env"
)

type MultilineMode uint8

const (
	MultilineEnv MultilineMode = iota
	MultilineDisabled
	MultilineEnabled
)

type Config struct {
	Level       Level
	Format      string
	Environment string
	Multiline   MultilineMode

	Output io.Writer
}

func (cfg *Config) defaults() {
	if cfg.Level == LevelUnknown {
		cfg.Level = LevelFrom(env.Get("LOG_LEVEL", "debug"))
	}

	if cfg.Format == "" {
		cfg.Format = env.Get("LOG_FORMAT", "json")
	}

	if cfg.Environment == "" {
		cfg.Environment = env.Get("ENVIRONMENT", "dev")
	}

	if cfg.Multiline == MultilineEnv {
		multiline := env.Get("LOG_MULTILINE", false)
		if multiline {
			cfg.Multiline = MultilineEnabled
		} else {
			cfg.Multiline = MultilineDisabled
		}
	}

	if cfg.Output == nil {
		cfg.Output = os.Stdout
	}
}
