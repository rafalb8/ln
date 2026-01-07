package ln

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"path/filepath"
)

type textHandler struct {
	io.Writer
	json func([]Attr, bool) []byte
}

func TextHandler(w io.Writer, multiline bool) *textHandler {
	// Choose if json.marshal must be idented or not
	marshaler := RenderJSON

	if multiline {
		marshaler = func(a []Attr, nl bool) []byte {
			buf := &bytes.Buffer{}
			json.Indent(buf, RenderJSON(a, nl), "", "  ")
			return buf.Bytes()
		}
	}

	return &textHandler{w, marshaler}
}

func (h *textHandler) Handle(ctx context.Context, r *Record) error {
	b := h.json(r.Attrs, false)
	if len(b) <= 2 {
		b = nil
	}

	buf := &bytes.Buffer{}
	buf.WriteString(r.Time.Format("15:04:05.000"))
	buf.WriteByte(' ')
	buf.WriteString(r.Level.Color())
	buf.WriteByte(' ')
	if r.Caller != "" {
		buf.WriteString("\x1b[90m")
		buf.WriteString(filepath.Base(r.Caller))
		buf.WriteString("\x1b[0m")
		buf.WriteByte(' ')
	}
	buf.WriteString("\x1b[97m")
	buf.WriteString(r.Message)
	if len(b) > 0 {
		buf.WriteByte(' ')
		buf.Write(b)
	}
	buf.WriteString("\x1b[0m\n")

	_, err := io.Copy(h, buf)
	return err
}
