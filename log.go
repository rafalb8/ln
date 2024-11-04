package log

import (
	"context"
	"sync"
	"time"
)

var recordPool = sync.Pool{
	New: func() any { return &Record{} },
}

type Logger struct {
	handler handler
	config  Config
}

// New returns logger with default configuration
func New() Logger {
	return NewWithConfiguration(Config{})
}

// NewWithConfiguration returns logger with given configuration
func NewWithConfiguration(cfg Config) Logger {
	cfg.defaults()
	var handler handler

	if cfg.Format == "json" {
		handler = JSONHandler(cfg.Output)
	} else {
		handler = TextHandler(cfg.Output, *cfg.Multiline)
	}

	return Logger{
		handler: handler,
		config:  cfg,
	}
}

func (l *Logger) Log(ctx context.Context, level Level, msg string, attrs []Attr) {
	if *l.config.Level > level {
		return
	}

	r := recordPool.Get().(*Record)
	defer recordPool.Put(r)

	r.Time = time.Now()
	r.Level = level
	r.Message = msg
	r.setCaller()
	r.Attrs = append([]Attr{
		String("caller", r.Caller),
		String("enviroment", l.config.Enviroment),
	}, attrs...)

	l.handler.Handle(ctx, r)
}

// Logs with default logger at level TRACE
func (l *Logger) Trace(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelTrace, msg, attrs)
}

// Logs with default logger at level DEBUG
func (l *Logger) Debug(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelDebug, msg, attrs)
}

// Logs with default logger at level INFO
func (l *Logger) Info(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelInfo, msg, attrs)
}

// Logs with default logger at level WARN
func (l *Logger) Warn(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelWarn, msg, attrs)
}

// Logs with default logger at level ERROR
func (l *Logger) Error(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelError, msg, attrs)
}

// Logs with default logger at level FATAL
func (l *Logger) Fatal(msg string, attrs ...Attr) {
	l.Log(context.Background(), LevelFatal, msg, attrs)
	l.panic(msg, attrs)
}

// Logs with default logger at level TRACE with given context
func (l *Logger) TraceCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelTrace, msg, attrs)
}

// Logs with default logger at level DEBUG with given context
func (l *Logger) DebugCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelDebug, msg, attrs)
}

// Logs with default logger at level INFO with given context
func (l *Logger) InfoCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelInfo, msg, attrs)
}

// Logs with default logger at level WARN with given context
func (l *Logger) WarnCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelWarn, msg, attrs)
}

// Logs with default logger at level ERROR with given context
func (l *Logger) ErrorCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelError, msg, attrs)
}

// Logs with default logger at level FATAL with given context
func (l *Logger) FatalCtx(ctx context.Context, msg string, attrs ...Attr) {
	l.Log(ctx, LevelFatal, msg, attrs)
	l.panic(msg, attrs)
}

func (*Logger) panic(msg string, attrs []Attr) {
	for _, attr := range attrs {
		err, ok := attr.Value.(error)
		if !ok {
			continue
		}
		if attr.Key == "error" {
			msg = err.Error()
			break
		}
	}
	panic(msg)
}
