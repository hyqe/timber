package timber

import (
	"io"
	"os"
)

type Jack interface {
	Debug(v any)
	Error(v any)
	Alert(v any)
}

// NewJack creates a Jack that uses fmt.Fprint as its encoder.
func NewJack(opts ...option) Jack {
	j := defaultJack()
	for _, opt := range opts {
		opt(j)
	}
	return j
}

type jack struct {
	io.Writer
	Level
	Format Formatter
}

func defaultJack() *jack {
	return &jack{
		Level:  DEBUG,
		Writer: newSyncWriter(os.Stdout),
		Format: defaultFormatter(STATUS),
	}
}

func (j *jack) log(lvl Level, v any) {
	if !j.Is(lvl) {
		return
	}

	io.Copy(j, j.Format(newLog(lvl, v)))
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
		j.Writer = w
	}
}

func WithFormatter(f Formatter) option {
	return func(j *jack) {
		j.Format = f
	}
}
