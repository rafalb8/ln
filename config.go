package ln

import (
	"io"
	"os"

	"github.com/rafalb8/ln/util/env"
)

type Format string

const (
	Text Format = "text"
	JSON Format = "json"
)

type Switch uint8

const (
	None Switch = iota
	On
	Off
)

type Config struct {
	Level       Level
	Format      Format
	Environment string
	Multiline   Switch
	CallerDepth int

	Output io.Writer
}

func (cfg *Config) defaults() {
	if cfg.Level == 0 {
		cfg.Level = LevelFrom(env.Get("LOG_LEVEL", "debug"))
	}

	if cfg.Format == "" {
		cfg.Format = env.Get[Format]("LOG_FORMAT", "text")
	}

	if cfg.Environment == "" {
		cfg.Environment = env.Get("ENVIRONMENT", "dev")
	}

	if cfg.Multiline == None {
		if env.Get("LOG_MULTILINE", false) {
			cfg.Multiline = On
		} else {
			cfg.Multiline = Off
		}
	}

	if cfg.Output == nil {
		cfg.Output = os.Stdout
	}
}
