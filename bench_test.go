package ln_test

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/rafalb8/ln"
)

var (
	ctx   = context.Background()
	date  = time.Date(2000, 20, 20, 20, 20, 20, 20, time.UTC)
	attrs = []ln.Attr{
		ln.Group("group", ln.Bool("bool", true)),
		ln.Int("int", 64),
		ln.Uint("uint", uint(128)),
		ln.Float("float", 1.23),
		ln.String("string", "str"),
		ln.Time("time", date),
		ln.Duration("duration", time.Since(date)),
		ln.Err(errors.New("error")),
		ln.Any("any", struct{}{}),
	}
)

func BenchmarkJSON(b *testing.B) {
	jsonLogger := ln.NewWithConfiguration(ln.Config{
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
	textLogger := ln.NewWithConfiguration(ln.Config{
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
