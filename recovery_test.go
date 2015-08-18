// lovingly stolen from negroni

package mux

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecovery(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()

	rec := NewRecovery()
	rec.Logger = log.New(buff, "[mux] ", 0)

	n := NewRouter()
	// replace log for testing
	n.UseMiddleware(rec)
	n.PathPrefix("/").Handler(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		panic("here is a panic!")
	}))
	req, _ := http.NewRequest("GET", "http://localhost:3000/foobar", nil)
	n.ServeHTTP(recorder, req)
	expect(t, recorder.Code, http.StatusInternalServerError)
	refute(t, recorder.Body.Len(), 0)
	refute(t, len(buff.String()), 0)
}
