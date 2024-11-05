package log_test

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	log "github.com/rafalb8/ln"
)

var (
	ctx   = context.Background()
	date  = time.Date(2000, 20, 20, 20, 20, 20, 20, time.UTC)
	attrs = []log.Attr{
		log.Group("group", log.Bool("bool", true)),
		log.Int("int", 64),
		log.Uint("uint", uint(128)),
		log.Float("float", 1.23),
		log.String("string", "str"),
		log.Time("time", date),
		log.Duration("duration", time.Since(date)),
		log.Err(errors.New("error")),
		log.Any("any", struct{}{}),
	}
)

func BenchmarkJSON(b *testing.B) {
	jsonLogger := log.NewWithConfiguration(log.Config{
		Format: "json",
		Output: io.Discard,
	})

	b.Run("NoAttrs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsonLogger.InfoCtx(ctx, "Test message")
		}
	})

	b.Run("WithAttrs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsonLogger.InfoCtx(ctx, "Test message", attrs...)
		}
	})
}

func BenchmarkText(b *testing.B) {
	textLogger := log.NewWithConfiguration(log.Config{
		Format: "text",
		Output: io.Discard,
	})

	b.Run("NoAttrs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			textLogger.InfoCtx(ctx, "Test message")
		}
	})

	b.Run("WithAttrs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			textLogger.InfoCtx(ctx, "Test message", attrs...)
		}
	})
}
