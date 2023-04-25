package timber

import "time"

type Log struct {
	CreatedAt time.Time
	Level     Level
	Data      any
}

func newLog(lvl Level, v any) Log {
	return Log{
		CreatedAt: time.Now().UTC(),
		Level:     lvl,
		Data:      v,
	}
}
