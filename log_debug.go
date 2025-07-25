//go:build debug

package ln

import "context"

func (l *Logger) Trace(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelTrace, msg, attrs)
}

func (l *Logger) Debug(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelDebug, msg, attrs)
}

func (l *Logger) TraceCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelTrace, msg, attrs)
}

func (l *Logger) DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelDebug, msg, attrs)
}

func Trace(msg string, attrs ...Attr) {
	Default.Trace(msg, attrs...)
}

func Debug(msg string, attrs ...Attr) {
	Default.Debug(msg, attrs...)
}

func TraceCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.TraceCtx(ctx, msg, attrs...)
}

func DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.DebugCtx(ctx, msg, attrs...)
}
