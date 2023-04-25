package timber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Lumber is the smallest unit in he timeber package, and is used to define a single log.
type Lumber interface {

	// CreatedAt returns the time the log was created.
	CreatedAt() time.Time

	// Returns the log level.
	Level() Level

	// Chop returns the actual log data. It provides a means of lazy
	// encoding/formatting. generally the level is evaluated first
	// before Chop() is call.
	Chop() io.Reader
}

type wood struct {
	ts   time.Time
	lvl  Level
	chop func() io.Reader
}

func (w *wood) Level() Level {
	return w.lvl
}

func (w *wood) Chop() io.Reader {
	return w.chop()
}

func (w *wood) CreatedAt() time.Time {
	return w.ts
}

func JSON(lvl Level, v any) Lumber {
	return &wood{
		ts:  time.Now().UTC(),
		lvl: lvl,
		chop: func() io.Reader {
			var buff bytes.Buffer
			buff.WriteString(fmt.Sprintf("%v: ", lvl.String())) // TODO: implement templates here
			json.NewEncoder(&buff).Encode(v)
			return &buff
		},
	}
}

func FPRINT(lvl Level, v any) Lumber {
	return &wood{
		ts:  time.Now().UTC(),
		lvl: lvl,
		chop: func() io.Reader {
			var buff bytes.Buffer
			buff.WriteString(fmt.Sprintf("%v: ", lvl.String())) // TODO: implement templates here
			fmt.Fprint(&buff, v)
			return &buff
		},
	}
}
