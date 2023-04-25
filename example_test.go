package timber_test

import (
	"bytes"
	"fmt"

	"github.com/hyqe/timber"
)

func Example() {
	timber.Debug("Timber!!!")
	// Output:
	// DEBUG: Timber!!!
}

func ExampleJack() {
	var output bytes.Buffer

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

	jack := timber.NewJack(
		timber.WithLevel(timber.DEBUG),
		timber.WithFormatter(formatter),
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
}

func ExampleWithLevel() {
	var output bytes.Buffer

	// set the logging level to ALERT.
	jack := timber.NewJack(
		timber.WithLevel(timber.ALERT),
		timber.WithWriter(&output),
	)

	jack.Debug("this will be ignored")
	jack.Error("this will be ignored")
	jack.Alert("this will alert")

	fmt.Println(output.String())
	// Output:
	// ALERT: this will alert
}

func ExampleWithFormatter() {
	var output bytes.Buffer

	formatter := func(l timber.Log) string {
		return fmt.Sprintf("my custom log: %v", l.Message)
	}

	jack := timber.NewJack(
		timber.WithFormatter(formatter),
		timber.WithWriter(&output),
	)

	jack.Debug("timber!")

	fmt.Println(output.String())
	// Output:
	// my custom log: timber!
}
