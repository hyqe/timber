package timber

import (
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	CreatedAt time.Time `json:"createdAt"`
	Level     Level     `json:"level"`
	Message   any       `json:"message"`
}

func newLog(lvl Level, v any) *Log {
	return &Log{
		CreatedAt: time.Now().UTC(),
		Level:     lvl,
		Message:   v,
	}
}

func (l *Log) Flat() map[string]string {
	return map[string]string{
		"CreatedAt": l.CreatedAt.Format(time.RFC3339),
		"Level":     l.Level.String(),
		"Message":   fmt.Sprint(l.Message),
	}
}

func (l *Log) MarshalJSON() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Log) JSON() []byte {
	b, _ := json.Marshal(l)
	return b
}
