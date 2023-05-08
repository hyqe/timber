# Timber 🪓
 
A flexible logging package.

```bash
go get github.com/hyqe/timber
```

Using the default logger.

```Go
timber.Debug("Timber!!!")
// Output:
// DEBUG: Timber!!!
```

Creating a custom logging instance.

```Go
formatter := func(l timber.Log) string {
	switch l.Level {
	case timber.DEBUG:
		return fmt.Sprintf("🚧: %v", l.Message)
	case timber.ERROR:
		return fmt.Sprintf("💩: %v", l.Message)
	default:
		return fmt.Sprintf("🤔: %v", l.Message)
	}
}

// create a custom timber.Jack 🪓
timber.Default = timber.NewJack(
	// set log levels
	timber.WithLevel(timber.DEBUG),
	// set custom Emitters and Custom Formatter
	timber.WithEmitters(timber.Console(formatter)),
)

timber.Debug("this will debug")
timber.Error("this will error")
timber.Alert("this will alert")
// Output:
// 🚧: this will debug
// 💩: this will error
// 🤔: this will alert
```