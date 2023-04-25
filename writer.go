package timber

import (
	"io"
	"sync"
)

func newSyncWriter(w io.Writer) io.Writer {
	return &muWriter{
		out: w,
	}
}

// muWriter is a wrapper around an io.Writer that writes with a mutex
type muWriter struct {
	sync.Mutex
	out io.Writer
}

// Write implements io.Writer and is safe for concurrency.
func (m *muWriter) Write(p []byte) (n int, err error) {
	m.Lock()
	defer m.Unlock()
	return m.out.Write(p)
}
