package common

import (
	"net/http"
)

type ResponseWriter interface {
	http.ResponseWriter
	Status() int
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) Status() int {
	return rw.status
}
