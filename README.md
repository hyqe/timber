# Timber 🪓
 
A flexible logging package.



Using the default logger.

```Go
timber.Debug("Timber!!!")
// Output:
// DEBUG: Timber!!!
```

Creating a logging instance, with a custom writer.

```Go
var output bytes.Buffer

// full control with custom formatters
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
jack := timber.NewJack(
    // set log levels
	timber.WithLevel(timber.DEBUG),

    // set custom formatter
	timber.WithFormatter(formatter),

    // set custom output
	timber.WithWriter(&output),
)

jack.Debug("this will debug")
jack.Error("this will error")
jack.Alert("this will alert")

fmt.Println(output.String())
// Output:
// 🚧: this will debug
// 💩: this will error
// 🤔: this will alert
```