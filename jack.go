package timber

import (
	"io"
	"os"
	"sync"
)

type Jack interface {
	Debug(v any)
	Error(v any)
	Alert(v any)
}

// NewJack creates a Jack that uses fmt.Fprint as its encoder.
func NewJack(opts ...option) Jack {
	return newJack().apply(opts...)
}

type jack struct {
	sync.Mutex
	Emitters []Emitter
	Level
}

func newJack() *jack {
	return &jack{
		Level: DEBUG,
		Emitters: []Emitter{
			FPRINTLN(os.Stdout, LEVEL),
		},
	}
}

func (j *jack) apply(opts ...option) *jack {
	j.Lock()
	defer j.Unlock()
	for _, opt := range opts {
		opt(j)
	}
	return j
}

func (j *jack) log(lvl Level, v any) {
	if !j.Is(lvl) {
		return
	}
	l := newLog(lvl, v)
	j.Lock()
	defer j.Unlock()
	emit(l, j.Emitters...)
}

func (j *jack) Debug(v any) {
	j.log(DEBUG, v)
}

func (j *jack) Error(v any) {
	j.log(ERROR, v)
}

func (j *jack) Alert(v any) {
	j.log(ALERT, v)
}

type option func(*jack)

// WithLevel sets the log level, which can change which logs gets emitted.
func WithLevel(lvl Level) option {
	return func(j *jack) {
		j.Level = lvl
	}
}

// AddPrinter adds a printer to the Emitters.
func AddPrinter(w io.Writer, f Formatter) option {
	return AddEmitters(FPRINTLN(w, f))
}

// SetPrinter sets/overwrites the Emitters with a printer.
func SetPrinter(w io.Writer, f Formatter) option {
	return SetEmitters(FPRINTLN(w, f))
}

// AddEmitters adds logging emitters.
func AddEmitters(fns ...Emitter) option {
	return func(j *jack) {
		j.Emitters = append(j.Emitters, fns...)
	}
}

// SetEmitters sets/overwrites the logging emitters.
func SetEmitters(fns ...Emitter) option {
	return func(j *jack) {
		j.Emitters = fns
	}
}
