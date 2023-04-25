package timber

import "io"

type Jack interface {
	Debug(v any)
	Error(v any)
	Alert(v any)
}

func NewJack(opts ...Option) Jack {
	var cfg Config
	for _, opt := range opts {
		opt(&cfg)
	}
	return &fprintJack{
		Picker: NewPicker(cfg.lvl, cfg.out),
	}
}

func NewJsonJack(opts ...Option) Jack {
	var cfg Config
	for _, opt := range opts {
		opt(&cfg)
	}
	return &jsonJack{
		Picker: NewPicker(cfg.lvl, cfg.out),
	}
}

type Option func(*Config)

type Config struct {
	lvl Level
	out io.Writer
}

func WithLevel(lvl Level) Option {
	return func(c *Config) {
		c.lvl = lvl
	}
}

func WithWriter(w io.Writer) Option {
	return func(c *Config) {
		c.out = w
	}
}

type fprintJack struct {
	Picker
}

func (j *fprintJack) Debug(v any) {
	j.Pick(FPRINT(DEBUG, v))
}
func (j *fprintJack) Alert(v any) {
	j.Pick(FPRINT(ALERT, v))
}
func (j *fprintJack) Error(v any) {
	j.Pick(FPRINT(ERROR, v))
}

type jsonJack struct {
	Picker
}

func (j *jsonJack) Debug(v any) {
	j.Pick(JSON(DEBUG, v))
}
func (j *jsonJack) Alert(v any) {
	j.Pick(JSON(ALERT, v))
}
func (j *jsonJack) Error(v any) {
	j.Pick(JSON(ERROR, v))
}
