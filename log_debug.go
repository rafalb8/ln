//go:build debug || !release

package ln

import "context"

func (l *Logger) Debug(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelDebug, msg, attrs)
}

func (l *Logger) DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelDebug, msg, attrs)
}

func Debug(msg string, attrs ...Attr) {
	Default.Debug(msg, attrs...)
}

func DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.DebugCtx(ctx, msg, attrs...)
}
