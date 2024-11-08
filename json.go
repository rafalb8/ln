package log

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"sync"
	"time"
)

type jsonHandler struct {
	io.Writer
}

func JSONHandler(w io.Writer) *jsonHandler {
	return &jsonHandler{w}
}

func (h *jsonHandler) Handle(ctx context.Context, r *Record) error {
	attrs := append([]Attr{
		Time("time", r.Time),
		String("level", r.Level.String()),
		String("msg", r.Message),
	}, r.Attrs...)

	_, err := h.Write(RenderJSON(attrs))
	return err
}

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func RenderJSON(attrs []Attr) []byte {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()

	renderAttrs(buf, attrs)
	buf.WriteByte('\n')

	return buf.Bytes()
}

func renderAttrs(buf *bytes.Buffer, a []Attr) {
	buf.WriteByte('{')
	for i, attr := range a {
		if i != 0 {
			buf.WriteByte(',')
		}

		buf.WriteByte('"')
		buf.WriteString(attr.Key)
		buf.WriteString(`":`)

		renderValue(buf, attr.Value)
	}
	buf.WriteByte('}')
}

func renderValue(buf *bytes.Buffer, val any) {
	switch val := val.(type) {
	case []Attr:
		renderAttrs(buf, val)
	case bool:
		buf.Write(strconv.AppendBool(buf.AvailableBuffer(), val))
	case int64:
		buf.Write(strconv.AppendInt(buf.AvailableBuffer(), val, 10))
	case uint64:
		buf.Write(strconv.AppendUint(buf.AvailableBuffer(), val, 10))
	case float64:
		buf.Write(strconv.AppendFloat(buf.AvailableBuffer(), val, 'e', -1, 64))
	case string:
		buf.Write(strconv.AppendQuoteToGraphic(buf.AvailableBuffer(), val))
	case time.Time:
		buf.WriteByte('"')
		buf.Write(val.AppendFormat(buf.AvailableBuffer(), time.RFC3339))
		buf.WriteByte('"')
	case time.Duration:
		buf.Write(strconv.AppendQuote(buf.AvailableBuffer(), val.String()))
	case error:
		buf.Write(strconv.AppendQuote(buf.AvailableBuffer(), val.Error()))
	case json.RawMessage:
		buf.Write(val)
	default:
		err := json.NewEncoder(buf).Encode(val)
		if err != nil {
			buf.Write(strconv.AppendQuote(buf.AvailableBuffer(), err.Error()))
		}
	}
}
