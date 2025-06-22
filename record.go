package ln

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
	runtime.Callers(4, pcs)

	frames := runtime.CallersFrames(pcs)
	frame, more := frames.Next()

	// If called from default.go, move to next frame
	if frame.File == defaultPath && more {
		frame, _ = frames.Next()
	}

	r.Caller = fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
