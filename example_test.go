package timber_test

import (
	"bytes"
	"fmt"

	"github.com/hyqe/timber"
)

func ExampleJack() {
	var buff bytes.Buffer

	jack := timber.NewJack(
		timber.WithWriter(&buff),
	)

	jack.Debug("this will debug")
	jack.Error("this will error")
	jack.Alert("this will alert")

	fmt.Println(buff.String())
	// Output:
	// DEBUG: this will debug
	// ERROR: this will error
	// ALERT: this will alert
}

func ExampleWithLevel() {
	var buff bytes.Buffer

	// set the logging level to ALERT.
	jack := timber.NewJack(
		timber.WithLevel(timber.ALERT),
		timber.WithWriter(&buff),
	)

	jack.Debug("this will be ignored")
	jack.Error("this will be ignored")
	jack.Alert("this will alert")

	fmt.Println(buff.String())
	// Output:
	// ALERT: this will alert
}
