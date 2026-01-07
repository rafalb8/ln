package ln

import (
	"bytes"
	"context"
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

	buf := &bytes.Buffer{}
	buf.WriteString(r.Time.Format(time.RFC3339))
	buf.WriteByte(' ')
	buf.WriteString(r.Level.String())

	if r.Caller != "" {
		buf.WriteString(" [")
		buf.WriteString(filepath.Base(r.Caller))
		buf.WriteByte(']')
	}

	buf.WriteString(" -- ")
	buf.WriteString(r.Message)
	if len(b) > 0 {
		buf.WriteByte(' ')
		buf.Write(b)
	}
	buf.WriteByte('\n')

	_, err := io.Copy(h, buf)
	return err
}
