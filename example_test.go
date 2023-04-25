package timber_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/hyqe/timber"
)

func Example() {
	timber.Debug("Timber!!!")
	// Output:
	// DEBUG: Timber!!!
}

func ExampleJack() {
	var output bytes.Buffer

	// full control with custom formatters
	formatter := func(l *timber.Log) string {
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

		// set custom printer
		timber.SetPrinter(&output, formatter),
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
		timber.SetPrinter(&output, timber.LEVEL),
	)

	jack.Debug("this will be ignored")
	jack.Error("this will be ignored")
	jack.Alert("this will alert")

	fmt.Println(output.String())
	// Output:
	// ALERT: this will alert
}

func ExampleCustomPrinter() {
	var output bytes.Buffer

	formatter := func(l *timber.Log) string {
		return fmt.Sprintf("my custom log: %v", l.Message)
	}

	jack := timber.NewJack(
		timber.SetPrinter(&output, formatter),
	)

	jack.Debug("timber!")

	fmt.Println(output.String())
	// Output:
	// my custom log: timber!
}

func ExampleHttpLog() {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	l := timber.NewHttpLog(r, http.StatusOK)

	timber.Debug(l)
	// Output:
	// DEBUG: GET / 200
}

func ExampleMiddleware() {
	middleware := timber.NewMiddleware(timber.NewJack())

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware(next)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)
	// Output:
	// DEBUG: GET / 200
}
