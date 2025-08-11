package middleware

import (
	"net/http"

	"github.com/rafalb8/ln"
)

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
	Size       int

	wroteHeader bool
}

func (w *responseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}

	w.wroteHeader = true
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(data []byte) (int, error) {
	n, err := w.ResponseWriter.Write(data)
	w.Size += n
	return n, err
}

func HTTP(l ln.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				l.ErrorCtx(r.Context(), "http panic", ln.Any("error", err))
			}
		}()

		writer := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(writer, r)

		l.InfoCtx(r.Context(), r.Method+" "+r.URL.Path,
			ln.Int("status", writer.StatusCode),
			ln.Int("size", writer.Size),
		)
	})
}
