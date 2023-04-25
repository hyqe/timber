package timber

import (
	"fmt"
	"io"
	"sync"
)

// Emitter is the last function called to process a log.
type Emitter func(*Log)

func FPRINTLN(w io.Writer, format Formatter) Emitter {
	var mu sync.Mutex
	return func(l *Log) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintln(w, format(l))
	}
}

func emit(l *Log, fns ...Emitter) {
	for _, fn := range fns {
		fn(l)
	}
}
