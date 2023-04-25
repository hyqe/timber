package timber

import (
	"fmt"
	"net/http"
)

// NewMiddleware creates a logger which maps http status to log level.
func NewMiddleware(j Jack) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := newResponseWriter(w)
			next.ServeHTTP(rw, r)

			l := NewHttpLog(r, rw.status)

			switch switchHttpStatus(rw.status) {
			case ERROR:
				j.Error(l)
			case DEBUG:
				j.Debug(l)
			}
		})
	}
}

func switchHttpStatus(status int) Level {
	switch {
	case status >= 500:
		return ERROR
	default:
		return DEBUG
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func NewHttpLog(r *http.Request, status int) *HttpLog {
	return &HttpLog{
		Method: r.Method,
		Path:   r.URL.Path,
		Status: status,
	}
}

type HttpLog struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Status int    `json:"status"`
}

func (r *HttpLog) String() string {
	return fmt.Sprintf("%v %v %v", r.Method, r.Path, r.Status)
}
