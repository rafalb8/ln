package ln

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
)

const (
	textTimeFormat = "15:04:05.000"

	// format: <time> <level> <caller> <message> <attrs>
	logFormatStr = "%s %s \x1b[90m%s\x1b[0m \x1b[97m%s %s\x1b[0m\n"
)

type textHandler struct {
	io.Writer
	json func([]Attr, bool) []byte
}

func TextHandler(w io.Writer, multiline MultilineMode) *textHandler {
	// Choose if json.marshal must be idented or not
	marshaler := RenderJSON

	if multiline == MultilineEnabled {
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

	_, err := fmt.Fprintf(h, logFormatStr,
		r.Time.Format(textTimeFormat),
		r.Level.Color(),
		filepath.Base(r.Caller),
		r.Message,
		string(b),
	)
	return err
}
