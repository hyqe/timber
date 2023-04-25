package timber

import "net/http"

// NewHttpStatusLogger creates a logger which maps http status to log level.
func NewHttpStatusLogger(j Jack) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := newResponseWriter(w)
			next.ServeHTTP(rw, r)
			switch switchHttpStatus(rw.status) {
			case ERROR:
				j.Error(r)
			case DEBUG:
				j.Debug(r)
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
