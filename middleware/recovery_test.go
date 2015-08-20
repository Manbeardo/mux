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

func TestRecovery(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()

	rec := NewRecovery()
	rec.Logger = log.New(buff, "[mux] ", 0)

	n := mux.NewRouter()
	// replace log for testing
	n.UseMiddleware(rec)
	n.PathPrefix("/").Handler(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		panic("here is a panic!")
	}))
	req, _ := http.NewRequest("GET", "http://localhost:3000/foobar", nil)
	n.ServeHTTP(recorder, req)
	assert.Equal(t, recorder.Code, http.StatusInternalServerError)
	assert.NotEqual(t, recorder.Body.Len(), 0)
	assert.NotEqual(t, len(buff.String()), 0)
}
