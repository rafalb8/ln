package log

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type handler interface {
	Handle(context.Context, *Record) error
}

type Record struct {
	Time    time.Time
	Level   Level
	Message string
	Attrs   []Attr

	Caller string
}

func (r *Record) setCaller() {
	// Get callers
	pcs := make([]uintptr, 2)
	n := runtime.Callers(5, pcs)
	pcs = pcs[:n]

	frames := runtime.CallersFrames(pcs)
	frame, _ := frames.Next()

	// If called from default.go, move to next frame
	if frame.File == defaultPath {
		frame, _ = frames.Next()
	}

	r.Caller = fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
