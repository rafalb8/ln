package log

import (
	"context"
	"runtime"
)

var (
	defaultLog  Logger = New()
	defaultPath        = func() string {
		pcs := make([]uintptr, 1)
		n := runtime.Callers(1, pcs)
		pcs = pcs[:n]
		frame, _ := runtime.CallersFrames(pcs).Next()
		return frame.File
	}()
)

func SetDefault(cfg Config) {
	defaultLog = NewWithConfiguration(cfg)
}

func Log(level Level, msg string, attrs ...Attr) {
	defaultLog.Log(context.Background(), level, msg, attrs)
}

func Trace(msg string, attrs ...Attr) {
	defaultLog.Trace(msg, attrs...)
}

func Debug(msg string, attrs ...Attr) {
	defaultLog.Debug(msg, attrs...)
}

func Info(msg string, attrs ...Attr) {
	defaultLog.Info(msg, attrs...)
}

func Warn(msg string, attrs ...Attr) {
	defaultLog.Warn(msg, attrs...)
}

func Error(msg string, attrs ...Attr) {
	defaultLog.Error(msg, attrs...)
}

func Fatal(msg string, attrs ...Attr) {
	defaultLog.Fatal(msg, attrs...)
}

func LogCtx(ctx context.Context, level Level, msg string, attrs ...Attr) {
	defaultLog.Log(ctx, level, msg, attrs)
}

func TraceCtx(ctx context.Context, msg string, attrs ...Attr) {
	defaultLog.TraceCtx(ctx, msg, attrs...)
}

func DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	defaultLog.DebugCtx(ctx, msg, attrs...)
}

func InfoCtx(ctx context.Context, msg string, attrs ...Attr) {
	defaultLog.InfoCtx(ctx, msg, attrs...)
}

func WarnCtx(ctx context.Context, msg string, attrs ...Attr) {
	defaultLog.WarnCtx(ctx, msg, attrs...)
}

func ErrorCtx(ctx context.Context, msg string, attrs ...Attr) {
	defaultLog.ErrorCtx(ctx, msg, attrs...)
}

func FatalCtx(ctx context.Context, msg string, attrs ...Attr) {
	defaultLog.FatalCtx(ctx, msg, attrs...)
}
