package timber

// Default enables using jack in the Global scope, as is
// often desired for logging packages.
var Default = NewDefaultJack()

func NewDefaultJack() *Jack {
	return NewJack(
		WithLevel(DEBUG),
		WithEmitters(Console(AsText())),
	)
}

// Alert creates new log with its level set to ALERT.
func Alert(v any) {
	Default.Alert(v)
}

// Alertf creates new log with its level set to ALERT.
func Alertf(format string, v any) {
	Default.Alertf(format, v)
}

// Error creates new log with its level set to ERROR.
func Error(v any) {
	Default.Error(v)
}

// Errorf creates new log with its level set to ERROR.
func Errorf(format string, v any) {
	Default.Errorf(format, v)
}

// Debug creates new log with its level set to DEBUG.
func Debug(v any) {
	Default.Debug(v)
}

// Debugf creates new log with its level set to DEBUG.
func Debugf(format string, v any) {
	Default.Debugf(format, v)
}
