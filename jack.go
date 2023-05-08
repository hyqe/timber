package timber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

// NewJack creates a Jack that uses fmt.Fprint as its encoder.
func NewJack(opts ...Option) *Jack {
	j := &Jack{
		Level:    DEBUG,
		Emitters: make([]Emitter, 0),
	}
	return j.Apply(opts...)
}

type Jack struct {
	sync.Mutex
	Level
	Emitters []Emitter
}

// Sync calls lock/unlock for you.
func (j *Jack) Sync(do func(*Jack)) {
	j.Lock()
	defer j.Unlock()
	do(j)
}

// Apply options to Jack.
func (j *Jack) Apply(opts ...Option) *Jack {
	for _, opt := range opts {
		opt(j)
	}
	return j
}

// Emit contains all the logic for admiting a log to all the emitters.
func (j *Jack) Emit(l Log) {
	j.Sync(func(j *Jack) {
		if len(j.Emitters) == 0 {
			return
		}
		if !j.Level.GTE(l.Level) {
			return
		}
		var wg sync.WaitGroup
		for _, emitter := range j.Emitters {
			wg.Add(1)
			go func(emit Emitter) {
				defer wg.Done()
				if emit == nil {
					return
				}
				emit(l)
			}(emitter)
		}
		wg.Wait()
	})
}

// Alert creates new log with its level set to ALERT.
func (j *Jack) Alert(v any) {
	j.Emit(Log{
		Level:   ALERT,
		Message: v,
	})
}

// Alertf creates new log with its level set to ALERT.
func (j *Jack) Alertf(format string, v any) {
	j.Alert(fmt.Sprintf(format, v))
}

// Error creates new log with its level set to ERROR.
func (j *Jack) Error(v any) {
	j.Emit(Log{
		Level:   ERROR,
		Message: v,
	})
}

// Errorf creates new log with its level set to ERROR.
func (j *Jack) Errorf(format string, v any) {
	j.Error(fmt.Sprintf(format, v))
}

// Debug creates new log with its level set to DEBUG.
func (j *Jack) Debug(v any) {
	j.Emit(Log{
		Level:   DEBUG,
		Message: v,
	})
}

// Debugf creates new log with its level set to DEBUG.
func (j *Jack) Debugf(format string, v any) {
	j.Debug(fmt.Sprintf(format, v))
}

// Option is any function that modifies a Jack.
type Option func(j *Jack)

// AddEmitters adds emitters to the existing list of emitters.
func AddEmitters(emitters ...Emitter) Option {
	return func(j *Jack) {
		j.Sync(func(j *Jack) {
			j.Emitters = append(j.Emitters, emitters...)
		})
	}
}

// WithEmitters removes all previous emitters and sets jack with the specified emitters.
func WithEmitters(emitters ...Emitter) Option {
	return func(j *Jack) {
		j.Sync(func(j *Jack) {
			j.Emitters = emitters
		})
	}
}

func WithLevel(lvl Level) Option {
	return func(j *Jack) {
		j.Sync(func(j *Jack) {
			j.Level = lvl
		})
	}
}

type Log struct {
	Level
	Message any
}

// Emitter is responsible for sending a log to its destination.
// this could be to the console, or somewhere else.
type Emitter = func(Log)

// Console emits logs to the console.
func Console(format Formatter) Emitter {
	return func(l Log) {
		io.Copy(os.Stdout, bytes.NewReader([]byte(format(l)+"\n")))
	}
}

// Formatter detirmins how a log will be formatted.
type Formatter func(l Log) string

// AsText formats the as "<level>: <message>"
func AsText() Formatter {
	return func(l Log) string {
		return fmt.Sprintf("%v: %v", l.Level, l.Message)
	}
}

// AsJson marshalls the entire log as json.
func AsJson() Formatter {
	return func(l Log) string {
		b, err := json.Marshal(l)
		if err != nil {
			return fmt.Sprintf("timber failed to marshal log as json: %v", b)
		}
		return string(b)
	}
}
