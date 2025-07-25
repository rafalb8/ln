//go:build !debug

package ln

import "context"

func (l *Logger) Trace(msg string, attrs ...Attr) {}

func (l *Logger) Debug(msg string, attrs ...Attr) {}

func (l *Logger) TraceCtx(ctx context.Context, msg string, attrs ...Attr) {}

func (l *Logger) DebugCtx(ctx context.Context, msg string, attrs ...Attr) {}

func Trace(msg string, attrs ...Attr) {}

func Debug(msg string, attrs ...Attr) {}

func TraceCtx(ctx context.Context, msg string, attrs ...Attr) {}

func DebugCtx(ctx context.Context, msg string, attrs ...Attr) {}
