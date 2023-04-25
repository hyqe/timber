package timber

var defaultJack *jack

func init() {
	// Go Examples is breaking when using the init() func.
	// apparently "go test ./..." replaces the os.Stdout
	// file after init is run, which causes it to not capture
	// the output correctly.
	// defaultJack = newJack()
}

func Apply(opts ...option) {
	if defaultJack == nil {
		defaultJack = newJack()
	}
	defaultJack.apply(opts...)
}

func Debug(v any) {
	if defaultJack == nil {
		defaultJack = newJack()
	}
	defaultJack.Debug(v)
}

func Error(v any, opts ...option) {
	if defaultJack == nil {
		defaultJack = newJack()
	}
	defaultJack.Error(v)
}

func Alert(v any, opts ...option) {
	if defaultJack == nil {
		defaultJack = newJack()
	}
	defaultJack.Alert(v)
}
