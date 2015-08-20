package middleware

import (
	"net/http"
	"time"
)

// Timeout middleware panics if the request hasn't completed after the time limit
type Timeout struct {
	Limit time.Duration
}

func (m *Timeout) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	doneChan := make(chan bool)
	go func() {
		next(w, r)
		doneChan <- true
	}()
	select {
	case <-doneChan:
	case <-time.After(m.Limit):
		panic("request processing exceeded time limit")
	}
}
