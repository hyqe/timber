package timber

import (
	"fmt"
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
	Out io.Writer
	Level
	Format Formatter
}

func newJack() *jack {
	return &jack{
		Level:  DEBUG,
		Out:    os.Stdout,
		Format: TEMPLATE(STATUS),
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
	j.Lock()
	defer j.Unlock()
	fmt.Fprintln(j.Out, j.Format(newLog(lvl, v)))
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

func WithLevel(lvl Level) option {
	return func(j *jack) {
		j.Level = lvl
	}
}

func WithWriter(w io.Writer) option {
	return func(j *jack) {
		j.Out = w
	}
}

func WithFormatter(f Formatter) option {
	return func(j *jack) {
		j.Format = f
	}
}
