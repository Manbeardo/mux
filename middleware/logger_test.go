// lovingly stolen from negroni

package middleware

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Manbeardo/mux"
)

func Test_Logger(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()

	l := NewLogger()
	l.Logger = log.New(buff, "[negroni] ", 0)

	n := mux.NewRouter()
	// replace log for testing
	n.UseMiddleware(l)
	n.UseMiddlewareHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))

	req, err := http.NewRequest("GET", "http://localhost:3000/foobar", nil)
	if err != nil {
		t.Error(err)
	}

	n.ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Code, http.StatusNotFound)
	assert.NotEqual(t, len(buff.String()), 0)
}
