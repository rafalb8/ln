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

func (r *Record) setCaller(skip int) {
	// Get callers
	pcs := make([]uintptr, 1)
	runtime.Callers(5+skip, pcs)

	// Extract caller frame
	frame, _ := runtime.CallersFrames(pcs).Next()
	r.Caller = fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
