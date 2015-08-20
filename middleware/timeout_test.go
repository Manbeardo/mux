package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type DelayHandler struct {
	Delay time.Duration
}

func (d DelayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	<-time.After(d.Delay)
}

func TestTimeoutPanics(t *testing.T) {
	m := Timeout{20 * time.Millisecond}

	defer func() {
		assert.NotNil(t, recover())
	}()

	m.ServeHTTP(httptest.NewRecorder(), nil, DelayHandler{30 * time.Millisecond}.ServeHTTP)
}

func TestTimeoutDoesntPanic(t *testing.T) {
	m := Timeout{20 * time.Millisecond}

	defer func() {
		assert.Nil(t, recover())
	}()

	m.ServeHTTP(httptest.NewRecorder(), nil, DelayHandler{1 * time.Microsecond}.ServeHTTP)
}

func TestTimeoutDoesStuff(t *testing.T) {
	m := Timeout{10 * time.Millisecond}

	touched := false

	m.ServeHTTP(httptest.NewRecorder(), nil, func(w http.ResponseWriter, r *http.Request) {
		touched = true
	})

	assert.True(t, touched)
}
