package log

import (
	"context"
	"runtime"
)

var (
	Default     Logger = New()
	defaultPath        = func() string {
		pcs := make([]uintptr, 1)
		n := runtime.Callers(1, pcs)
		pcs = pcs[:n]
		frame, _ := runtime.CallersFrames(pcs).Next()
		return frame.File
	}()
)

func Log(level Level, msg string, attrs ...Attr) {
	Default.Log(context.Background(), level, msg, attrs)
}

func Trace(msg string, attrs ...Attr) {
	Default.Trace(msg, attrs...)
}

func Debug(msg string, attrs ...Attr) {
	Default.Debug(msg, attrs...)
}

func Info(msg string, attrs ...Attr) {
	Default.Info(msg, attrs...)
}

func Warn(msg string, attrs ...Attr) {
	Default.Warn(msg, attrs...)
}

func Error(msg string, attrs ...Attr) {
	Default.Error(msg, attrs...)
}

func Fatal(msg string, attrs ...Attr) {
	Default.Fatal(msg, attrs...)
}

func LogCtx(ctx context.Context, level Level, msg string, attrs ...Attr) {
	Default.Log(ctx, level, msg, attrs)
}

func TraceCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.TraceCtx(ctx, msg, attrs...)
}

func DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.DebugCtx(ctx, msg, attrs...)
}

func InfoCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.InfoCtx(ctx, msg, attrs...)
}

func WarnCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.WarnCtx(ctx, msg, attrs...)
}

func ErrorCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.ErrorCtx(ctx, msg, attrs...)
}

func FatalCtx(ctx context.Context, msg string, attrs ...Attr) {
	Default.FatalCtx(ctx, msg, attrs...)
}
