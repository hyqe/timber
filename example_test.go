package timber_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/hyqe/timber"
)

func ExampleJack() {
	var output bytes.Buffer

	jack := timber.NewJack(
		timber.WithWriter(&output),
	)

	jack.Debug("this will debug")
	jack.Error("this will error")
	jack.Alert("this will alert")

	fmt.Println(output.String())
	// Output:
	// DEBUG: this will debug
	// ERROR: this will error
	// ALERT: this will alert
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

	formatter := func(l timber.Log) io.Reader {
		return strings.NewReader(fmt.Sprintf("my custom log: %v", l.Message))
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
