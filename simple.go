package ln

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"
)

type simpleHandler struct {
	io.Writer
}

func SimpleHandler(w io.Writer) *simpleHandler {
	return &simpleHandler{w}
}

func (h *simpleHandler) Handle(ctx context.Context, r *Record) error {
	attrs := make([]Attr, 0, len(r.Attrs))
	for _, a := range r.Attrs {
		if a.Key == "caller" || a.Key == "environment" {
			continue
		}
		attrs = append(attrs, a)
	}

	b := RenderJSON(attrs, false)
	if len(b) <= 2 {
		b = nil
	}

	// Format: <time> <level> [<caller>] -- <message> <attrs>
	_, err := fmt.Fprintf(h, "%s %s [%s] -- %s %s\n",
		r.Time.Format(time.RFC3339),
		r.Level.String(),
		filepath.Base(r.Caller),
		r.Message,
		string(b),
	)
	return err
}
