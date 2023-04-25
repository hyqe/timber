package timber_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyqe/timber"
)

func TestNewMiddlewareJack_StatusInternalServerError(t *testing.T) {

	emit := func(l *timber.Log) {
		if l.Level != timber.ERROR {
			t.Fatal(l.Level)
		}
	}

	jack := timber.NewJack(timber.SetEmitters(emit))
	middleware := timber.NewMiddlewareJack(jack)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	handler := middleware(next)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)
}

func TestNewMiddlewareJack_StatusBadRequest(t *testing.T) {

	emit := func(l *timber.Log) {
		if l.Level != timber.DEBUG {
			t.Fatal(l.Level)
		}
	}

	jack := timber.NewJack(timber.SetEmitters(emit))
	middleware := timber.NewMiddlewareJack(jack)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	handler := middleware(next)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)
}
